package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/stub", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("./composite-data.json")
		if err != nil {
			panic(err)
		}

		var data RespBody
		err = json.NewDecoder(f).Decode(&data)
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
