package redis

import (
	"context"
	"log"

	"api/src/config"

	"github.com/go-redis/redis/v8"
)

func New() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Conf.Redis.URL,
		Password: config.Conf.Redis.Password,
		DB:       0,
	})

	context := context.Background()

	_, err := client.Ping(context).Result()

	if err != nil {
		log.Fatal("failed to connect redis", err)
	}

	log.Print("success to connect redis!")

	return client
}
