package initialize

import (
	"blog/global"
	"blog/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func inItRedis() {
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort),
		Password: utils.RedisPassword,
		DB:       11,
	})
	_, err := global.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic("redis ping error")
	}
}
