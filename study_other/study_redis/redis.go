package study_redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var client *redis.Client

// Set 方法 设置key和value，以及key的过期时间expiration 返回值：error
func Set(ctx context.Context, key string, val interface{}) error {
	expiration := 30 * time.Second
	return client.Set(ctx, key, val, expiration).Err()
}

// Get 方法 返回值和错误信息
func Get(ctx context.Context, k string) (string, error) {
	return client.Get(ctx, k).Result()
}

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "106.55.186.222:6379",
		Password: "asdfgqwet123df12345", // no password set
		DB:       0,                     // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("redis ping err: " + err.Error())
	}
}
