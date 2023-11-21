package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/internal/server"
	"github.com/izaakdale/vancouver-conditions-backend/stub"
)

func main() {
	stub.Run()

	opt, err := redis.ParseURL(os.Getenv("UPSTASH_REDIS_URL"))
	if err != nil {
		panic(err)
	}
	cli := redis.NewClient(opt)

	server.Start(cli)
}
