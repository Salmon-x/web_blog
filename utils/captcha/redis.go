package captcha

import (
	"blog/global"
	"context"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"time"
)

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "SALMON_CAPTCHA_",
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

// 实现store接口

func (rs *RedisStore) Set(id string, value string) error {
	err := global.RedisClient.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.RedisClient.Get(rs.Context, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		err := global.RedisClient.Del(rs.Context, key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
