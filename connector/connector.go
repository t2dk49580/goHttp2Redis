package connector

import (
	"github.com/go-redis/redis"
)

func Get(redisaddr string, redispass string, rediskey string) (string, error) {
	client := redis.NewClient(&redis.Options{Addr: redisaddr, Password: redispass, DB: 0})
	return client.Do("get", rediskey).String()
}

func Set(redisaddr string, redispass string, rediskey string, redisvalue string) (string, error) {
	client := redis.NewClient(&redis.Options{Addr: redisaddr, Password: redispass, DB: 0})
	return client.Do("set", rediskey, redisvalue).String()
}
