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

func Get (key string) interface{} {
	value, err := redisdb.Get(key).Result()
	if err != nil {
		fmt.Println("Get redis value failed,err:", err)
		return nil
	}
	return value
}

func Set (key string, value interface{}) {
	err := redisdb.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("Set redis failed,err:", err)
		return
	}
	return
}

func Del (keys ...string) {
	for _, key := range keys {
		err := redisdb.Del(key).Err()
		if err != nil {
			fmt.Println("Redis del failed,err:", err)
			return
		}
	}
	return
}

func HMSet (key string, fields map[string]interface{}) {
	err := redisdb.HMSet(key, fields).Err()
	if err != nil {
		fmt.Println("HMSet failed,err:", err)
		return
	}
	return
}

func HMGet (key string, fields ...string) []interface{} {
	var result []interface{}
	for _, field := range fields {
		tmp, err := redisdb.HMGet(key, field).Result()
		if err != nil || tmp[0] == nil{
			fmt.Println("HMGet failed or nil,err:", err)
			return nil
		}
		result = append(result, tmp[0])
	}
	return result
}

func HDel (key string, fields ...string) {
	for _, field := range fields {
		err := redisdb.HDel(key, field).Err()
		if err != nil {
			fmt.Println("HDel failed,err:", err)
			return
		}
	}
	return
}