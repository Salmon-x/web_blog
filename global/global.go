package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	RedisClient *redis.Client
	Db          *gorm.DB
)
