package main

type ResortData struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Country             string  `json:"country"`
	Continent           string  `json:"continent"`
	CurrentBaseTempC    float64 `json:"current_temp_c"`
	SnowFall1Day        float64 `json:"snow_fall_1_day"`
	SnowFall3Days       float64 `json:"snow_fall_3_days"`
	SnowFall5Days       float64 `json:"snow_fall_5_days"`
	PrecipitationStatus string  `json:"precipitation_status"`
}

type RespBody struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Continent string `json:"continent"`
	Forecast  []struct {
		Date     string  `json:"date"`
		Time     string  `json:"time"`
		PrecipMm float64 `json:"precip_mm"`
		PrecipIn float64 `json:"precip_in"`
		RainMm   float64 `json:"rain_mm"`
		RainIn   float64 `json:"rain_in"`
		SnowMm   float64 `json:"snow_mm"`
		SnowIn   float64 `json:"snow_in"`
		Base     struct {
			FreshsnowCm float64 `json:"freshsnow_cm"`
			FreshsnowIn float64 `json:"freshsnow_in"`
			TempC       float64 `json:"temp_c"`
			TempF       float64 `json:"temp_f"`
			FeelslikeC  float64 `json:"feelslike_c"`
			FeelslikeF  float64 `json:"feelslike_f"`
		} `json:"base"`
		Mid struct {
			FreshsnowCm float64 `json:"freshsnow_cm"`
			FreshsnowIn float64 `json:"freshsnow_in"`
			TempC       any     `json:"temp_c"`
			TempF       any     `json:"temp_f"`
			FeelslikeC  any     `json:"feelslike_c"`
			FeelslikeF  any     `json:"feelslike_f"`
		} `json:"mid"`
		Upper struct {
			FreshsnowCm float64 `json:"freshsnow_cm"`
			FreshsnowIn float64 `json:"freshsnow_in"`
			TempC       any     `json:"temp_c"`
			TempF       any     `json:"temp_f"`
			FeelslikeC  any     `json:"feelslike_c"`
			FeelslikeF  any     `json:"feelslike_f"`
		} `json:"upper"`
	} `json:"forecast"`
}

type FullRespBody struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Continent string `json:"continent"`
	Forecast  []struct {
		Date          string  `json:"date"`
		Time          string  `json:"time"`
		LowcloudPct   float64 `json:"lowcloud_pct"`
		MidcloudPct   float64 `json:"midcloud_pct"`
		HighcloudPct  float64 `json:"highcloud_pct"`
		TotalcloudPct float64 `json:"totalcloud_pct"`
		FrzglvlFt     float64 `json:"frzglvl_ft"`
		FrzglvlM      float64 `json:"frzglvl_m"`
		PrecipMm      float64 `json:"precip_mm"`
		PrecipIn      float64 `json:"precip_in"`
		RainMm        float64 `json:"rain_mm"`
		RainIn        float64 `json:"rain_in"`
		SnowMm        float64 `json:"snow_mm"`
		SnowIn        float64 `json:"snow_in"`
		HumPct        float64 `json:"hum_pct"`
		DewpointC     float64 `json:"dewpoint_c"`
		DewpointF     float64 `json:"dewpoint_f"`
		VisKm         float64 `json:"vis_km"`
		VisMi         float64 `json:"vis_mi"`
		SlpMb         float64 `json:"slp_mb"`
		SlpIn         float64 `json:"slp_in"`
		Base          struct {
			WxDesc         string  `json:"wx_desc"`
			WxCode         int     `json:"wx_code"`
			WxIcon         string  `json:"wx_icon"`
			FreshsnowCm    float64 `json:"freshsnow_cm"`
			FreshsnowIn    float64 `json:"freshsnow_in"`
			TempC          float64 `json:"temp_c"`
			TempF          float64 `json:"temp_f"`
			FeelslikeC     float64 `json:"feelslike_c"`
			FeelslikeF     float64 `json:"feelslike_f"`
			WinddirDeg     float64 `json:"winddir_deg"`
			WinddirCompass string  `json:"winddir_compass"`
			WindspdMph     float64 `json:"windspd_mph"`
			WindspdKmh     float64 `json:"windspd_kmh"`
			WindspdKts     float64 `json:"windspd_kts"`
			WindspdMs      float64 `json:"windspd_ms"`
			WindgstMph     float64 `json:"windgst_mph"`
			WindgstKmh     float64 `json:"windgst_kmh"`
			WindgstKts     float64 `json:"windgst_kts"`
			WindgstMs      float64 `json:"windgst_ms"`
		} `json:"base"`
		Mid struct {
			WxDesc         string  `json:"wx_desc"`
			WxCode         int     `json:"wx_code"`
			WxIcon         string  `json:"wx_icon"`
			FreshsnowCm    float64 `json:"freshsnow_cm"`
			FreshsnowIn    float64 `json:"freshsnow_in"`
			TempC          any     `json:"temp_c"`
			TempF          any     `json:"temp_f"`
			FeelslikeC     any     `json:"feelslike_c"`
			FeelslikeF     any     `json:"feelslike_f"`
			WinddirDeg     float64 `json:"winddir_deg"`
			WinddirCompass string  `json:"winddir_compass"`
			WindspdMph     float64 `json:"windspd_mph"`
			WindspdKmh     float64 `json:"windspd_kmh"`
			WindspdKts     float64 `json:"windspd_kts"`
			WindspdMs      float64 `json:"windspd_ms"`
			WindgstMph     any     `json:"windgst_mph"`
			WindgstKmh     any     `json:"windgst_kmh"`
			WindgstKts     any     `json:"windgst_kts"`
			WindgstMs      any     `json:"windgst_ms"`
		} `json:"mid"`
		Upper struct {
			WxDesc         string  `json:"wx_desc"`
			WxCode         int     `json:"wx_code"`
			WxIcon         string  `json:"wx_icon"`
			FreshsnowCm    float64 `json:"freshsnow_cm"`
			FreshsnowIn    float64 `json:"freshsnow_in"`
			TempC          float64 `json:"temp_c"`
			TempF          float64 `json:"temp_f"`
			FeelslikeC     float64 `json:"feelslike_c"`
			FeelslikeF     float64 `json:"feelslike_f"`
			WinddirDeg     float64 `json:"winddir_deg"`
			WinddirCompass string  `json:"winddir_compass"`
			WindspdMph     float64 `json:"windspd_mph"`
			WindspdKmh     float64 `json:"windspd_kmh"`
			WindspdKts     float64 `json:"windspd_kts"`
			WindspdMs      float64 `json:"windspd_ms"`
			WindgstMph     any     `json:"windgst_mph"`
			WindgstKmh     any     `json:"windgst_kmh"`
			WindgstKts     any     `json:"windgst_kts"`
			WindgstMs      any     `json:"windgst_ms"`
		} `json:"upper"`
	} `json:"forecast"`
}
