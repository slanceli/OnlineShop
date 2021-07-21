package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dbcon = "root:admin@tcp(127.0.0.1:3306)/?charset=utf8"

func DBInit() { //初始化数据库
	_db, err := sql.Open("mysql", dbcon)
	DB = _db
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		panic(err)
	}
}