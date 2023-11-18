package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	t, err := time.LoadLocation("America/Vancouver")
	if err != nil {
		panic(err)
	}

	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	cli := redis.NewClient(opt)

	sch := gocron.NewScheduler(t)
	// sch.Cron("0 */6 * * *").Do(func() {
	sch.Cron("* * * * *").Do(func() {
		fmt.Println("Cron says hello")

		bytes, err := os.ReadFile("./composite-data.json")
		if err != nil {
			panic(err)
		}

		cli.Set("latest-conditions", bytes, 0)
	})
	sch.StartAsync()

	mux.HandleFunc("/stub", func(w http.ResponseWriter, r *http.Request) {
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

	http.ListenAndServe("localhost:8080", cors.Default().Handler(mux))
}
