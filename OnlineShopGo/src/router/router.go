package router

import (
	"OnlineShopGo/src/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

//验证用户信息中间件
func AuthenticateUserInfo () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//初始化路由
func InitRouter () {
	Router = gin.Default()
	//store := cookie.NewStore([]byte(utils.CookiePassword))
	//Router.Use(sessions.Sessions("OnlieShopSession", store))
	Router.GET("/login", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})


	Router.POST("/login", func(c *gin.Context) {
		userName := c.PostForm("name")
		userPasswd := c.PostForm("passwd")
		if user.Login(userName, userPasswd) {
			c.String(http.StatusOK, "successful")
		} else {
			c.String(http.StatusOK, "failed")
		}
	})
}

func Run (address string) {
	err := Router.Run(address)
	if err != nil {
		fmt.Println("Run Router falied, err:", err)
	}
}