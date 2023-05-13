package study_redis

import (
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func Setex(key string, value interface{}, expiration time.Duration) error {
	return client.Set(key, value, expiration).Err()
}

// Get 获取 key 对应的 value
func Get(key string) (string, error) {
	return client.Get(key).Result()

}

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "106.55.186.222:6379",
		Password: "remember", // no password set
		DB:       0,          // use default DB
	})
}
