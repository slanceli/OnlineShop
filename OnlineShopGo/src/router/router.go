package router

import (
	"OnlineShopGo/src/user"
	"OnlineShopGo/src/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

//验证用户登录信息中间件
func AuthenticateUserInfo () gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("login")
		if v == nil {
			c.String(http.StatusUnauthorized, "未登录")
			c.Abort()
		} else {
			fmt.Println(v)
			c.Next()
		}
	}
}

//初始化路由
func InitRouter () {
	Router = gin.Default()
	store := cookie.NewStore([]byte(utils.CookiePassword))
	Router.Use(sessions.Sessions("online_shop", store))
	Router.GET("/login", AuthenticateUserInfo(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	Router.GET("/register", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})


	Router.POST("/login", func(c *gin.Context) {
		userName := c.PostForm("name")
		userPasswd := c.PostForm("passwd")
		if user.Login(userName, userPasswd) {
			session := sessions.Default(c)
			session.Set("login", 1)
			err := session.Save()
			if err != nil {
				fmt.Println("Save session failed, err:", err)
			}
			c.String(http.StatusOK, "successful")
		} else {
			c.String(http.StatusOK, "failed")
		}
	})
	Router.POST("/register", func(c *gin.Context) {
		userName := c.PostForm("name")
		userPasswd := c.PostForm("passwd")
		c.String(http.StatusOK, user.Register(userName, userPasswd))
	})
}

func Run (address string) {
	err := Router.Run(address)
	if err != nil {
		fmt.Println("Run Router falied, err:", err)
	}
}