package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	Host     string
	Port     int
	User     string
	Pwd      string
	Database string
}

var DB *sqlx.DB

func init() {
	mysql := &Mysql{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Pwd:      "12345678",
		Database: "zhuxiaozhuan",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", mysql.User, mysql.Pwd, mysql.Host, mysql.Port, mysql.Database)

	database, err := sqlx.Open("mysql", dsn)

	if err != nil {
		fmt.Println("open mysql err", err)
		return
	}
	DB = database
	//defer database.Close()
}
