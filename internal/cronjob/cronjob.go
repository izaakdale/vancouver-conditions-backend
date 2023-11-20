package cronjob

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/pkg/api"
)

func StartAsync() error {
	t, err := time.LoadLocation("America/Vancouver")
	if err != nil {
		return err
	}
	sch := gocron.NewScheduler(t)
	sch.Cron("0 */6 * * *").Do(func() {
		fmt.Println("Cron says hello")

		fmt.Println(os.Getenv("UPSTASH_REDIS_URL"))
		chronOpts, err := redis.ParseURL(os.Getenv("UPSTASH_REDIS_URL"))
		if err != nil {
			log.Printf("error trying to connect to redis\n")
			return
		}
		chronCli := redis.NewClient(chronOpts)

		rec := api.Record{
			Data: []api.FullBody{},
		}

		// wg := sync.WaitGroup{}
		// mu := sync.Mutex{}

		weatherApiEndpoint := os.Getenv("WEATHER_API_ENDPOINT")
		apiKey := os.Getenv("WEATHER_API_KEY")

		// wg.Add(len(searchParams))

		for searchQuery, adds := range searchParams {
			s := searchQuery
			a := adds
			// go func() {
			u := fmt.Sprintf("%s/%s?unitGroup=metric&key=%s&contentType=json", weatherApiEndpoint, s, apiKey)

			req, err := http.NewRequest(http.MethodGet, u, nil)
			if err != nil {
				log.Printf("error creating request for %s: %+v\n", s, err)
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("error when fetching data from %s: %+v\n", s, err)
				return
			}
			var fb api.FullBody
			err = json.NewDecoder(resp.Body).Decode(&fb)
			if err != nil {
				log.Printf("error decoding response from %s: %+v\n", s, err)
				return
			}

			fb.ForecastUrl = a.forecastUrl
			fb.WebCamUrl = a.webCamUrl
			fb.Title = a.title
			fb.GoogleMapsUrl = a.googleMapsUrl

			// mu.Lock()
			rec.Data = append(rec.Data, fb)
			// mu.Unlock()

			// wg.Done()
			// }()
		}

		// wg.Wait()
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
