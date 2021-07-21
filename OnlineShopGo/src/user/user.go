package user

import (
	"OnlineShopGo/src/dao"
	"database/sql"
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