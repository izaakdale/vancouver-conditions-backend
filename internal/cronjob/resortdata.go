package cronjob

// search params maps the api search query to any additional data
// I want included in the record which is not supplied by the api
var searchParams = map[string]additionalData{
	"whistler-blackcomb-mountain": {
		title:         "Whistler Blackcomb",
		webCamUrl:     "https://www.whistlerblackcomb.com/the-mountain/mountain-conditions/mountain-cams.aspx",
		forecastUrl:   "https://www.snow-forecast.com/resorts/Whistler-Blackcomb/6day/mid",
		googleMapsUrl: "https://maps.app.goo.gl/7YTvXnCQPS32mxE9A",
	},
	"mt-baker-washington": {
		title:         "Mount Baker",
		webCamUrl:     "https://www.snowstash.com/usa/washington/mt-baker/snow-cams",
		forecastUrl:   "https://www.snow-forecast.com/resorts/Mount-Baker/6day/mid",
		googleMapsUrl: "https://maps.app.goo.gl/gaqSji8YiTb8RacY6",
	},
	"20955-hemlock-valley-rd": {
		title:         "Sasquatch Mountain Resort",
		webCamUrl:     "https://sasquatchmountain.ca/weather-and-conditions/webcams/",
		forecastUrl:   "https://www.snow-forecast.com/resorts/HemlockResort/6day/mid",
		googleMapsUrl: "https://maps.app.goo.gl/o5CWVongU85nwqhT7",
	},
	"cypress-mountain-vancouver": {
		title:         "Cypress Mountain",
		webCamUrl:     "https://cypressmountain.com/downhill-conditions-and-cams",
		forecastUrl:   "https://www.snow-forecast.com/resorts/Cypress-Mountain/6day/mid",
		googleMapsUrl: "https://maps.app.goo.gl/pJkSrmDLMb4RikAd8",
	},
	// omitted due to weather reports being almost identical to cypress.
	// "seymour-mountain-vancouver": {
	// 	webCamUrl:   "https://www.youtube.com/watch?v=vLawo-FrBKk",
	// 	forecastUrl: "https://www.snow-forecast.com/resorts/Mount-Seymour/6day/mid",
	// },
	// "grouse-mountain-vancouver": {
	// 	webCamUrl:   "https://www.grousemountain.com/web-cams",
	// 	forecastUrl: "https://www.snow-forecast.com/resorts/Grouse-Mountain/6day/mid",
	// },
}

type additionalData struct {
	title         string
	webCamUrl     string
	forecastUrl   string
	googleMapsUrl string
}
