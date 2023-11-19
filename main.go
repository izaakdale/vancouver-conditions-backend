package main

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/internal/cronjob"
	"github.com/izaakdale/vancouver-conditions-backend/internal/server"
	"github.com/izaakdale/vancouver-conditions-backend/stub"
)

func main() {
	stub.Run()

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	cli := redis.NewClient(opt)

	err = cronjob.StartAsync()
	if err != nil {
		panic(err)
	}
	server.Start(cli)
}
