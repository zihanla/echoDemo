package model

import (
	_ "github.com/go-sql-driver/mysql" // _ 表示只执行init函数
	"github.com/jmoiron/sqlx"
	"log"
)

// Db 定义一个操作数据的指针全局变量
var Db *sqlx.DB

// Init 初始化数据库
func init() {
	db, err := sqlx.Open(`mysql`, `root:root@tcp(127.0.0.1:3306)/news?charset=utf8mb4&parseTime=true`)
	if err != nil {
		// 打印致命错误 先输入log 再退出程序
		log.Fatalln(err.Error())
	}
	Db = db
}
