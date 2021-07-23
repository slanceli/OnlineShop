package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var redisdb *redis.Client

//初始化redis连接
func Init() {
	redisdb = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	ret, err := redisdb.Ping().Result()
	fmt.Println("Result of the redisPing:", ret)
	if err != nil {
		fmt.Println("Redis init failed,err:", err)
		panic(err)
	}
}

func Get (key string) string {
	value, err := redisdb.Get(key).Result()
	if err != nil {
		fmt.Println("Get redis value failed,err:", err)
		return ""
	}
	return value
}

func Set (key string, value interface{}) error {
	err := redisdb.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("Set redis failed,err:", err)
		return err
	}
	return nil
}