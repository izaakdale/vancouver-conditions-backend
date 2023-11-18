package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	"github.com/rs/cors"
)

func main() {
	runStub()

	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	cli := redis.NewClient(opt)

	mux := http.NewServeMux()
	mux.HandleFunc("/resort-data", func(w http.ResponseWriter, r *http.Request) {
		val, err := cli.Get("latest-conditions").Bytes()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var data RespBody
		err = json.Unmarshal(val, &data)
		if err != nil {
			panic(err)
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	t, err := time.LoadLocation("America/Vancouver")
	if err != nil {
		panic(err)
	}
	sch := gocron.NewScheduler(t)
	// sch.Cron("0 */6 * * *").Do(func() {
	sch.Cron("* * * * *").Do(func() {
		fmt.Println("Cron says hello")

		chronOpts, err := redis.ParseURL("redis://localhost:6379")
		if err != nil {
			log.Printf("error trying to connect to redis\n")
			return
		}
		chronCli := redis.NewClient(chronOpts)

		rb := RespBody{
			Data: []ResortReport{},
		}

		wg := sync.WaitGroup{}
		endpoints := []string{"stub/whistler", "stub/seymour"}
		wg.Add(len(endpoints))

		for _, e := range endpoints {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:9090/"+e, nil)
			if err != nil {
				panic(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}
			var rr ResortReport
			json.NewDecoder(resp.Body).Decode(&rr)

			rb.Data = append(rb.Data, rr)
			wg.Done()
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

	http.ListenAndServe("localhost:8080", cors.Default().Handler(mux))
}

func runStub() {
	mux := http.NewServeMux()

	mux.HandleFunc("/stub/whistler", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./data-whistler.json")
		bytes, _ := io.ReadAll(f)

		var fb FullBody
		json.Unmarshal(bytes, &fb)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fb)
	})
	mux.HandleFunc("/stub/seymour", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./data-seymour.json")
		bytes, _ := io.ReadAll(f)

		var fb FullBody
		json.Unmarshal(bytes, &fb)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fb)
	})
	go http.ListenAndServe("localhost:9090", mux)
}
