package main

import (
	"OnlineShopGo/src/dao"
	"OnlineShopGo/src/router"
	"OnlineShopGo/src/utils"
)

func main () {
	dao.DBInit()
	router.InitRouter()
	router.Run(utils.LocalAddress)
}