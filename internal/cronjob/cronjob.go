package cronjob

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/pkg/api"
)

func StartAsync() {
	t, err := time.LoadLocation("America/Vancouver")
	if err != nil {
		panic(err)
	}
	sch := gocron.NewScheduler(t)
	sch.Cron("0 */6 * * *").Do(func() {
		fmt.Println("Cron says hello")

		chronOpts, err := redis.ParseURL("redis://localhost:6379")
		if err != nil {
			log.Printf("error trying to connect to redis\n")
			return
		}
		chronCli := redis.NewClient(chronOpts)

		rb := api.RespBody{
			Data: []api.ResortReport{},
		}

		wg := sync.WaitGroup{}
		mu := sync.Mutex{}

		paths := []string{"stub/whistler", "stub/seymour", "stub/cypress"}
		wg.Add(len(paths))

		for _, path := range paths {
			p := path
			go func() {
				req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/"+p, nil)
				if err != nil {
					panic(err)
				}
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					panic(err)
				}
				var rr api.ResortReport
				json.NewDecoder(resp.Body).Decode(&rr)

				mu.Lock()
				rb.Data = append(rb.Data, rr)
				mu.Unlock()

				wg.Done()
			}()
		}

		wg.Wait()
		bytes, err := json.Marshal(rb)
		if err != nil {
			panic(err)
		}

		err = chronCli.Set("latest-conditions", bytes, 0).Err()
		if err != nil {
			panic(err)
		}
	})
	sch.StartAsync()
}
