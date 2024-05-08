package kpl

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Mysql struct {
	Host     string
	Port     int
	User     string
	Pwd      string
	Database string
}

type GameHero struct {
	Id       int    `db:"id"`
	Cname    string `db:"cname"`
	Title    string `db:"title"`
	NewType  int    `db:"new_type"`
	HeroType int    `db:"hero_type"`
	SkinName string `db:"skin_name"`
	HeroImg  string `db:"hero_img"`
}

func skin() {
	mysql := &Mysql{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Pwd:      "12345678",
		Database: "zhuxiaozhuan",
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", mysql.User, mysql.Pwd, mysql.Host, mysql.Port, mysql.Database)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}

	defer db.Close()

	var hero []GameHero

	err = db.Select(&hero, "SELECT * from game_hero id=105")
	if err != nil {
		return
	}

	for _, gameHero := range hero {
		skinName := strings.Split(gameHero.SkinName, "|")
		for i, _ := range skinName {
			err := download(gameHero.Id, i+1, "/Users/zhangcheng/WWW/kpl/skin/")
			if err != nil {
				return
			}
		}
	}
}
