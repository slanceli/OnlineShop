package main

import (
	"OnlineShopGo/src/dao"
	"OnlineShopGo/src/redis"
	"OnlineShopGo/src/router"
	"OnlineShopGo/src/utils"
)

func main () {
	dao.DBInit()
	redis.Init()
	router.InitRouter()
	router.Run(utils.LocalAddress)
}