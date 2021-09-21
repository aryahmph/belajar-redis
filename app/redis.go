package app

import (
	"belajar-redis/helper"
	"context"
	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Username: "aryahmph",
		Password: "aryahmph",
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		helper.PanicIfError(err)
	}
	return rdb
}
