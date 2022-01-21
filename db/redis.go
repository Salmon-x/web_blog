package db

import (
	"blog/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort),
		Password: utils.RedisPassword,
		DB: 7,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("redis ping error")
	}
}
