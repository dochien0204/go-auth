package configs

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.TODO()

func RedisDatabase() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
