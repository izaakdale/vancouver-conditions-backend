package main

import (
	"github.com/go-redis/redis"
	"github.com/izaakdale/vancouver-conditions-backend/internal/cronjob"
	"github.com/izaakdale/vancouver-conditions-backend/internal/server"
	"github.com/izaakdale/vancouver-conditions-backend/stub"
)

func main() {
	stub.Run()

	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	cli := redis.NewClient(opt)

	cronjob.StartAsync()
	server.Start(cli)
}
