package goods

import (
	"OnlineShopGo/src/dao"
	"OnlineShopGo/src/jsonify"
	"OnlineShopGo/src/redis"
	"database/sql"
	"fmt"
)

type Goods struct {
	Name string
	Price float64
	Description string
	Goods_Left int
	Imgbase64 string
}

func AddGoods (goods Goods) bool {
	redis.HMSet(goods.Name, map[string]interface{}{
		"Price" : goods.Price,
		"Description" : goods.Description,
		"Goods_Left" : goods.Goods_Left,
		"Imgbase64" : goods.Imgbase64,
	})
	sqlStr := "INSERT INTO onlineshop.goods VALUES (NULL, ?, ?, ?, ?, ?)"
	ret, err := dao.DB.Exec(sqlStr, goods.Name, goods.Price, goods.Description, goods.Goods_Left, goods.Imgbase64)
	if err != nil {
		fmt.Println("Add goods failed, err:", err)
		return false
	}
	redis.HMSet(goods.Name, map[string]interface{}{
		"Price" : goods.Price,
		"Description" : goods.Description,
		"Goods_Left" : goods.Goods_Left,
		"Imgbase64" : goods.Imgbase64,
	})
	fmt.Println(ret)
	return true
}

func DeleteGoods (goodsName string) string {
	var sqlStr string
	var sqlResult string
	redisResult := redis.HMGet(goodsName, "Price")
	if redisResult == nil {
		sqlStr := "SELECT Price FROM onlineshop.goods WHERE Name = ?"
		err := dao.DB.QueryRow(sqlStr, goodsName).Scan(&sqlResult)
		if err == sql.ErrNoRows {
			return "notexist"
		}
		redis.HMSet(goodsName, map[string]interface{}{"id": sqlResult})
	}
	redis.HDel(goodsName, "id", "Name", "Price", "Description", "Goods_Left", "Imgbase64")
	redis.Del("AllGoods")
	sqlStr = "DELETE FROM onlineshop.goods WHERE Name = ?"
	ret, err := dao.DB.Exec(sqlStr, goodsName)
	if err != nil {
		fmt.Println("Delete goods failed, err:", err)
		return "failed"
	}
	redis.HDel(goodsName, "id", "Name", "Price", "Description", "Goods_Left", "Imgbase64")
	fmt.Println(ret)
	return "successful"
}

func GetGoods (num int) string {
	sqlStr := "SELECT Name, Goods_Left, Price, Description, Imgbase64 FROM onlineshop.goods LIMIT ?"
	rows, err := dao.DB.Query(sqlStr, num)
	if err != nil {
		fmt.Println("Get goods failed, err:", err)
		return "failed"
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			fmt.Println("rowsClose failed,err:", err)
		}
	}()
	return jsonify.Jsonify(rows)
}