package goods

import (
	"OnlineShopGo/src/dao"
	"database/sql"
	"fmt"
)

type Goods struct {
	Name string
	Price float64
	Description string
	Left int
	Imgbase64 string
}

func AddGoods (goods Goods) bool {
	sqlStr := "INSERT INTO onlineshop.goods VALUES (NULL, ?, ?, ?, ?, ?)"
	ret, err := dao.DB.Exec(sqlStr, goods.Name, goods.Price, goods.Description, goods.Left, goods.Imgbase64)
	if err != nil {
		fmt.Println("Add goods failed, err:", err)
		return false
	}
	fmt.Println(ret)
	return true
}

func DeleteGoods (goodsName string) string {
	sqlStr := "SELECT id FROM onlineshop.goods WHERE Name = ?"
	err := dao.DB.QueryRow(sqlStr, goodsName).Scan(nil)
	if err == sql.ErrNoRows {
		return "notexist"
	}
	sqlStr = "DELETE FROM onlineshop.goods WHERE Name = ?"
	ret, err := dao.DB.Exec(sqlStr, goodsName)
	if err != nil {
		fmt.Println("Delete goods failed, err:", err)
		return "failed"
	}
	fmt.Println(ret)
	return "successful"
}