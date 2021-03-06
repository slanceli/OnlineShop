package router

import (
	"OnlineShopGo/src/goods"
	order2 "OnlineShopGo/src/order"
	"OnlineShopGo/src/user"
	"OnlineShopGo/src/utils"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Router *gin.Engine

//验证用户登录信息中间件
func AuthenticateUserInfo () gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUrl := ""
		if len(c.Request.RequestURI) >= 6 {
			requestUrl = c.Request.RequestURI[1:6]
		}
		session := sessions.Default(c)
		v := session.Get("name")
		if v == nil {
			c.String(http.StatusUnauthorized, "未登录")
			c.Abort()
		} else {
			vStr := v.(string)
			if requestUrl == "admin" {
				if vStr != "admin" {
					c.String(http.StatusUnauthorized, "权限不足")
					c.Abort()
				} else {
					c.Next()
				}
			}
		}
	}
}

//初始化路由
func InitRouter () {
	Router = gin.Default()
	store := cookie.NewStore([]byte(utils.CookiePassword))
	Router.Use(sessions.Sessions("online_shop", store))
	admin := Router.Group("/admin", AuthenticateUserInfo())
	{
		admin.GET("", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
		admin.GET("/deletegoods", func(c *gin.Context) {
			goodsName := c.Query("GoodsName")
			if res := goods.DeleteGoods(goodsName); res == "successful"{
				c.String(http.StatusOK, "successful")
			} else {
				c.String(http.StatusInternalServerError, res)
			}
		})
		admin.GET("/updategoods", func(c *gin.Context) {
			Attributes := c.Query("Attributes")
			Value := c.Query("Value")
			GoodsName := c.Query("GoodsName")
			c.String(http.StatusOK, goods.UpdateGoods(Attributes, GoodsName, Value))
		})
		admin.GET("/updateorder", func(c *gin.Context) {
			OrderId := c.Query("OrderId")
			c.String(http.StatusOK, order2.UpdateOrder(OrderId))
		})
		admin.GET("/deleteorder", func(c *gin.Context) {
			OrderId := c.Query("OrderId")
			c.String(http.StatusOK, order2.DeleteOrder(OrderId))
		})
		admin.GET("/getallorder", func(c *gin.Context) {
			c.String(http.StatusOK, order2.GetAllOrder())
		})

		admin.POST("/addgoods", func(c *gin.Context) {
			body := goods.Goods{}
			if err := c.ShouldBindJSON(&body); err != nil {
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{"error": err.Error()})
					fmt.Println("BingJSON failed, error:", err)
				return
			}
			fmt.Println(body)
			if goods.AddGoods(body) {
				c.String(http.StatusOK, "successful")
			} else {
				c.String(http.StatusInternalServerError, "failed")
			}
		})
	}

	order := Router.Group("/order", AuthenticateUserInfo())
	{
		order.GET("", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
		order.GET("/getorder", func(c *gin.Context) {
			session := sessions.Default(c)
			v := session.Get("name")
			c.String(http.StatusOK, order2.GetOrder(v.(string)))
		})
		order.GET("/cancelorder", func(c *gin.Context) {
			OrderId := c.Query("OrderId")
			c.String(http.StatusOK, order2.CancelOrder(OrderId))
		})

		order.POST("/makeorder", func(c *gin.Context) {
			body := order2.Order{}
			if err := c.ShouldBindJSON(&body); err != nil {
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{"error": err.Error()})
				fmt.Println("BingJSON failed, error:", err)
				return
			}
			fmt.Println(body)
			session := sessions.Default(c)
			v := session.Get("name")
			body.OrderId = utils.GetGUID().Hex()
			body.UserName = v.(string)
			c.String(http.StatusOK, order2.MakeOrder(body))
		})
	}

	Router.GET("/getgoods", func(c *gin.Context) {
		num, _ := strconv.Atoi(c.Query("num"))
		c.String(http.StatusOK, goods.GetGoods(num))
	})
	Router.GET("/register", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	Router.POST("/login", func(c *gin.Context) {
		userName := c.PostForm("name")
		userPasswd := c.PostForm("passwd")
		if user.Login(userName, userPasswd) {
			session := sessions.Default(c)
			session.Set("name", userName)
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
	Router.POST("/changepasswd", AuthenticateUserInfo(), func(c *gin.Context) {
		oldPasswd := c.PostForm("oldPasswd")
		newPasswd := c.PostForm("newPasswd")
		session := sessions.Default(c)
		v := session.Get("name")
		fmt.Println(v)
		c.String(http.StatusOK, user.ChangePasswd(v.(string), oldPasswd, newPasswd))
	})
}

func Run (address string) {
	err := Router.Run(address)
	if err != nil {
		fmt.Println("Run Router failed, err:", err)
	}
}