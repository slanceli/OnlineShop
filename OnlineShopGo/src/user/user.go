package user

import (
	"OnlineShopGo/src/dao"
	"database/sql"
	"fmt"
)

func Login (name string, passwd string) bool {
	var sqlResult string
	sqlStr := "SELECT passwd FROM onlineshop.user WHERE name = ?"
	err := dao.DB.QueryRow(sqlStr, name).Scan(&sqlResult)
	if err == sql.ErrNoRows || sqlResult != passwd {
		return false
	}
	return true
}

func Register (name string, passwd string) (string) {
	sqlResult := ""
	sqlStr := "SELECT name FROM onlineshop.user WHERE name = ?"
	_ = dao.DB.QueryRow(sqlStr, name).Scan(&sqlResult)
	if sqlResult != "" {
		return "exist"
	}
	sqlStr = "INSERT INTO onlineshop.user VALUES (NULL, ?, ?)"
	ret, err := dao.DB.Exec(sqlStr, name, passwd)
	if err != nil {
		fmt.Println("Register failed, err:", err)
		return "failed"
	}
	fmt.Println(ret)
	return "successful"
}