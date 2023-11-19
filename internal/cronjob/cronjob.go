package cronjob

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/pkg/api"
)

var searchParams = map[string]additionalData{
	"whistler-blackcomb-mountain": {
		webCamUrl:   "https://www.whistlerblackcomb.com/the-mountain/mountain-conditions/mountain-cams.aspx",
		forecastUrl: "https://www.snow-forecast.com/resorts/Whistler-Blackcomb/6day/mid",
	},
	"seymour-mountain-vancouver": {
		webCamUrl:   "https://www.youtube.com/watch?v=vLawo-FrBKk",
		forecastUrl: "https://www.snow-forecast.com/resorts/Mount-Seymour/6day/mid",
	},
	"cypress-mountain-vancouver": {
		webCamUrl:   "https://cypressmountain.com/downhill-conditions-and-cams",
		forecastUrl: "https://www.snow-forecast.com/resorts/Cypress-Mountain/6day/mid",
	},
	// "grouse-mountain-vancouver": {
	// 	webCamUrl:   "https://www.grousemountain.com/web-cams",
	// 	forecastUrl: "https://www.snow-forecast.com/resorts/Grouse-Mountain/6day/mid",
	// },
}

type additionalData struct {
	webCamUrl   string
	forecastUrl string
}

func StartAsync() error {
	t, err := time.LoadLocation("America/Vancouver")
	if err != nil {
		return err
	}
	sch := gocron.NewScheduler(t)
	sch.Cron("* * * * *").Do(func() {
		fmt.Println("Cron says hello")

		chronOpts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			log.Printf("error trying to connect to redis\n")
			return
		}
		chronCli := redis.NewClient(chronOpts)

		rec := api.Record{
			Data: []api.FullBody{},
		}

		wg := sync.WaitGroup{}
		mu := sync.Mutex{}

		// weatherApiEndpoint := os.Getenv("WEATHER_API_ENDPOINT")
		// apiKey := os.Getenv("WEATHER_API_KEY")

		wg.Add(len(searchParams))

		for path, adds := range searchParams {
			p := path
			a := adds
			go func() {
				// u := fmt.Sprintf("%s/%s?unitGroup=metric&key=%s&contentType=json", weatherApiEndpoint, p, apiKey)
				u := fmt.Sprintf("%s/%s", "http://localhost:9090", p)

				req, err := http.NewRequest(http.MethodGet, u, nil)
				if err != nil {
					log.Printf("error creating request for %s: %+v\n", p, err)
					return
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("error when fetching data from %s: %+v\n", p, err)
					return
				}
				var fb api.FullBody
				err = json.NewDecoder(resp.Body).Decode(&fb)
				if err != nil {
					log.Printf("error decoding response from %s: %+v\n", p, err)
					return
				}

				fb.ForecastUrl = a.forecastUrl
				fb.WebCamUrl = a.webCamUrl

				mu.Lock()
				rec.Data = append(rec.Data, fb)
				mu.Unlock()

				wg.Done()
			}()
		}

		wg.Wait()
		bytes, err := json.Marshal(rec)
		if err != nil {
			log.Printf("error marshalling responses to bytes: %+v\n", err)
			return
		}

		err = chronCli.Set("latest-conditions", bytes, 0).Err()
		if err != nil {
			log.Printf("error setting data in redis: %+v\n", err)
			return
		}
	})
	sch.StartAsync()
	return nil
}
