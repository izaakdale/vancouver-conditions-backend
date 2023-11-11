package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	mux := http.NewServeMux()

	// allow origins
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
		// org := r.Header.Get("Origin")
		// return org == "http://localhost:5173"
	}}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		f, err := os.Open("./data.json")
		if err != nil {
			panic(err)
		}

		var data RespBody
		err = json.NewDecoder(f).Decode(&data)
		if err != nil {
			panic(err)
		}

		toSend := ResortData{
			ID:               data.ID,
			Name:             data.Name,
			Country:          data.Country,
			Continent:        data.Continent,
			CurrentBaseTempC: data.Forecast[0].Base.TempC,
		}

		if data.Forecast[0].RainMm == 0 && data.Forecast[0].SnowMm == 0 {
			toSend.PrecipitationStatus = "Still"
		}
		if data.Forecast[0].RainMm > data.Forecast[0].SnowMm {
			toSend.PrecipitationStatus = "Raining"
		} else {
			toSend.PrecipitationStatus = "Snowing"
		}

		var snow1, snow3, snow5 float64
		for i, f := range data.Forecast {
			if i < 8 {
				snow1 += f.Base.FreshsnowCm
			} else if i < 30 {
				snow3 += f.Base.FreshsnowCm
			} else {
				snow5 += f.Base.FreshsnowCm
			}
		}
		toSend.SnowFall1Day = snow1
		toSend.SnowFall3Days = snow1 + snow3
		toSend.SnowFall5Days = snow1 + snow3 + snow5

		var i int
		for {
			if err := conn.WriteJSON(toSend); err != nil {
				// probably due to client close, break to trigger deferred close of conn
				break
			}
			time.Sleep(time.Second * 10)
			i++
		}
	})

	http.ListenAndServe("localhost:8080", mux)
}
