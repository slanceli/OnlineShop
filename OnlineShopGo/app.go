package main

import (
	"OnlineShopGo/src/dao"
	"OnlineShopGo/src/redis"
	"OnlineShopGo/src/router"
	"OnlineShopGo/src/utils"
	"fmt"
)

func main () {
	dao.DBInit()
	redis.Init()
	_ = redis.Set("hello", 100)
	fmt.Println(redis.Get("hello"))
	router.InitRouter()
	router.Run(utils.LocalAddress)
}