run:
	WEATHER_API_ENDPOINT=https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/ \
	WEATHER_API_KEY=THCN6H3N9W42YRTRYGP66SLZJ \
	REDIS_URL=redis://localhost:6379 \
	HOST=192.168.1.31 \
	PORT=8080 \
	go run .