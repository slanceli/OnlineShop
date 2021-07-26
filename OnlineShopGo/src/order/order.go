package order

import (
	"OnlineShopGo/src/dao"
	"OnlineShopGo/src/redis"
	"fmt"
	"strconv"
	"time"
)

type Order struct {
	OrderId string
	UserName string
	GoodsName string
	Address string
	Num int
	Date string
	State string
}

func GetOrder () {

}

//创建订单并生成订单编号
func MakeOrder (order Order) string {
	var goods_left int
	sqlStr := "SELECT Goods_Left FROM onlineshop.goods WHERE Name = ?"
	err := dao.DB.QueryRow(sqlStr, order.GoodsName).Scan(&goods_left)
	if err != nil {
		fmt.Println("Query goods_left number failed,err:", err)
		return "failed"
	}
	redis.HMSet(order.GoodsName, map[string]interface{}{
		"Goods_Left" : strconv.Itoa(goods_left - order.Num),
	})
	sqlStr = "UPDATE onlineshop.goods SET Goods_Left = ? WHERE Name = ?"
	_, err = dao.DB.Exec(sqlStr, goods_left - order.Num, order.GoodsName)
	if err != nil {
		fmt.Println("Update goods_left failed,err:", err)
		return "failed"
	}
	defer func() {
		time.Sleep(time.Millisecond * 500)
		redis.HMSet(order.GoodsName, map[string]interface{}{
			"Goods_Left" : strconv.Itoa(goods_left - order.Num),
		})
	}()
	sqlStr = "INSERT INTO onlineshop.order VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err = dao.DB.Exec(sqlStr, order.OrderId, order.UserName, order.GoodsName, order.Address, order.Num, order.Date, 0)
	if err != nil {
		fmt.Println("Make order failed, err:", err)
		return "failed"
	}
	return "successful"
}

func DeleteOrder () {

}