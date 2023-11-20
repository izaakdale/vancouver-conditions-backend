package api

// response sent to the frontend
type RespBody struct {
	Data []ResortReport `json:"data"`
}

// db stored version of RespBody. Includes all data.
type Record struct {
	Data []FullBody `json:"data"`
}

// each resorts data within the frontend response
type ResortReport struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Description     string  `json:"description"`
	ResolvedAddress string  `json:"resolvedAddress"`
	WebCamUrl       string  `json:"webcamurl"`
	ForecastUrl     string  `json:"forecasturl"`
	Title           string  `json:"title"`
	Days            []struct {
		Datetime    string  `json:"datetime"`
		Precip      float64 `json:"precip"`
		Precipprob  float64 `json:"precipprob"`
		Precipcover float64 `json:"precipcover"`
		Preciptype  any     `json:"preciptype"`
		Snow        float64 `json:"snow"`
		Snowdepth   float64 `json:"snowdepth"`
	}
	CurrentConditions struct {
		Temp       float64 `json:"temp"`
		Feelslike  float64 `json:"feelslike"`
		Precip     any     `json:"precip"`
		Precipprob float64 `json:"precipprob"`
		Snow       float64 `json:"snow"`
		Snowdepth  float64 `json:"snowdepth"`
		Preciptype any     `json:"preciptype"`
		Visibility float64 `json:"visibility"`
		Conditions string  `json:"conditions"`
		Icon       string  `json:"icon"`
	}
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	Description string `json:"description"`
	Ends        string `json:"ends"`
	EndsEpoch   int    `json:"endsEpoch"`
	Event       string `json:"event"`
	Headline    string `json:"headline"`
	ID          string `json:"id"`
	Language    string `json:"language"`
	Link        string `json:"link"`
	Onset       string `json:"onset"`
	OnsetEpoch  int    `json:"onsetEpoch"`
}

// this is the response from the API
type FullBody struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Description     string  `json:"description"`
	// ADDED MYSELF REVISE
	WebCamUrl   string `json:"webcamurl"`
	ForecastUrl string `json:"forecasturl"`
	Title       string `json:"title"`
	//
	Days []struct {
		Datetime       string   `json:"datetime"`
		DatetimeEpoch  int      `json:"datetimeEpoch"`
		Tempmax        float64  `json:"tempmax"`
		Tempmin        float64  `json:"tempmin"`
		Temp           float64  `json:"temp"`
		Feelslikemax   float64  `json:"feelslikemax"`
		Feelslikemin   float64  `json:"feelslikemin"`
		Feelslike      float64  `json:"feelslike"`
		Dew            float64  `json:"dew"`
		Humidity       float64  `json:"humidity"`
		Precip         float64  `json:"precip"`
		Precipprob     float64  `json:"precipprob"`
		Precipcover    float64  `json:"precipcover"`
		Preciptype     any      `json:"preciptype"`
		Snow           float64  `json:"snow"`
		Snowdepth      float64  `json:"snowdepth"`
		Windgust       float64  `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        float64  `json:"winddir"`
		Pressure       float64  `json:"pressure"`
		Cloudcover     float64  `json:"cloudcover"`
		Visibility     float64  `json:"visibility"`
		Solarradiation float64  `json:"solarradiation"`
		Solarenergy    float64  `json:"solarenergy"`
		Uvindex        float64  `json:"uvindex"`
		Severerisk     float64  `json:"severerisk"`
		Sunrise        string   `json:"sunrise"`
		SunriseEpoch   int      `json:"sunriseEpoch"`
		Sunset         string   `json:"sunset"`
		SunsetEpoch    int      `json:"sunsetEpoch"`
		Moonphase      float64  `json:"moonphase"`
		Conditions     string   `json:"conditions"`
		Description    string   `json:"description"`
		Icon           string   `json:"icon"`
		Stations       []string `json:"stations"`
		Source         string   `json:"source"`
		Hours          []struct {
			Datetime       string   `json:"datetime"`
			DatetimeEpoch  int      `json:"datetimeEpoch"`
			Temp           float64  `json:"temp"`
			Feelslike      float64  `json:"feelslike"`
			Humidity       float64  `json:"humidity"`
			Dew            float64  `json:"dew"`
			Precip         float64  `json:"precip"`
			Precipprob     float64  `json:"precipprob"`
			Snow           float64  `json:"snow"`
			Snowdepth      float64  `json:"snowdepth"`
			Preciptype     any      `json:"preciptype"`
			Windgust       float64  `json:"windgust"`
			Windspeed      float64  `json:"windspeed"`
			Winddir        float64  `json:"winddir"`
			Pressure       float64  `json:"pressure"`
			Visibility     float64  `json:"visibility"`
			Cloudcover     float64  `json:"cloudcover"`
			Solarradiation float64  `json:"solarradiation"`
			Solarenergy    float64  `json:"solarenergy"`
			Uvindex        float64  `json:"uvindex"`
			Severerisk     float64  `json:"severerisk"`
			Conditions     string   `json:"conditions"`
			Icon           string   `json:"icon"`
			Stations       []string `json:"stations"`
			Source         string   `json:"source"`
		} `json:"hours"`
	} `json:"days"`
	Alerts   []any `json:"alerts"`
	Stations struct {
		Cwsk struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"CWSK"`
		Bhchk struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"BHCHK"`
		Cvod struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"CVOD"`
		Cwae struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"CWAE"`
		Bt001 struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"BT001"`
		F6229 struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"F6229"`
		Cwpn struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"CWPN"`
		Cwly struct {
			Distance     float64 `json:"distance"`
			Latitude     float64 `json:"latitude"`
			Longitude    float64 `json:"longitude"`
			UseCount     int     `json:"useCount"`
			ID           string  `json:"id"`
			Name         string  `json:"name"`
			Quality      int     `json:"quality"`
			Contribution float64 `json:"contribution"`
		} `json:"CWLY"`
	} `json:"stations"`
	CurrentConditions struct {
		Datetime       string   `json:"datetime"`
		DatetimeEpoch  int      `json:"datetimeEpoch"`
		Temp           float64  `json:"temp"`
		Feelslike      float64  `json:"feelslike"`
		Humidity       float64  `json:"humidity"`
		Dew            float64  `json:"dew"`
		Precip         any      `json:"precip"`
		Precipprob     float64  `json:"precipprob"`
		Snow           float64  `json:"snow"`
		Snowdepth      float64  `json:"snowdepth"`
		Preciptype     any      `json:"preciptype"`
		Windgust       any      `json:"windgust"`
		Windspeed      float64  `json:"windspeed"`
		Winddir        float64  `json:"winddir"`
		Pressure       float64  `json:"pressure"`
		Visibility     float64  `json:"visibility"`
		Cloudcover     float64  `json:"cloudcover"`
		Solarradiation float64  `json:"solarradiation"`
		Solarenergy    float64  `json:"solarenergy"`
		Uvindex        float64  `json:"uvindex"`
		Conditions     string   `json:"conditions"`
		Icon           string   `json:"icon"`
		Stations       []string `json:"stations"`
		Source         string   `json:"source"`
		Sunrise        string   `json:"sunrise"`
		SunriseEpoch   int      `json:"sunriseEpoch"`
		Sunset         string   `json:"sunset"`
		SunsetEpoch    int      `json:"sunsetEpoch"`
		Moonphase      float64  `json:"moonphase"`
	} `json:"currentConditions"`
}
