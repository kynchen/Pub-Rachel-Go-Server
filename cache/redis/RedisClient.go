package redis

import (
	"log"
	"time"
)

var redisClient = RedisClient
var redisContext = CTX

func Set(key string, value interface{}, expire int) bool {
	err := redisClient.Set(redisContext, key, value, time.Second*time.Duration(expire)).Err()
	if err != nil {
		log.Println("redis缓存异常", err)
		panic(err)
	}
	return true
}

func Get(key string) interface{} {
	result, err := redisClient.Get(redisContext, key).Result()
	if err != nil {
		log.Println("redis获取数据异常", err)
		panic(err)
	}
	return result
}
