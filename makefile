run:
	UPSTASH_REDIS_URL=redis://localhost:6379 \
	HOST=192.168.1.31 \
	PORT=8080 \
	go run .