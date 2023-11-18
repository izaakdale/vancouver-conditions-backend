package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/pkg/api"
	"github.com/rs/cors"
)

func Start(cli *redis.Client /*TODO interface*/) {

	mux := http.NewServeMux()
	mux.HandleFunc("/resort-data", func(w http.ResponseWriter, r *http.Request) {
		val, err := cli.Get("latest-conditions").Bytes()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var data api.RespBody
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

	http.ListenAndServe(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), cors.Default().Handler(mux))
}
