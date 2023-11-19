package stub

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/izaakdale/vancouver-conditions-backend/pkg/api"
)

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/whistler-blackcomb-mountain", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./data-whistler.json")
		bytes, _ := io.ReadAll(f)

		var fb api.FullBody
		json.Unmarshal(bytes, &fb)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fb)
	})
	mux.HandleFunc("/seymour-mountain-vancouver", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./data-seymour.json")
		bytes, _ := io.ReadAll(f)

		var fb api.FullBody
		json.Unmarshal(bytes, &fb)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fb)
	})
	mux.HandleFunc("/cypress-mountain-vancouver", func(w http.ResponseWriter, r *http.Request) {
		f, _ := os.Open("./data-cypress.json")
		bytes, _ := io.ReadAll(f)

		var fb api.FullBody
		json.Unmarshal(bytes, &fb)

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fb)
	})
	go http.ListenAndServe("localhost:9090", mux)
}
