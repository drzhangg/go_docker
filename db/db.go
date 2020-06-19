package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Engine *gorm.DB

var dataSource = "root:root@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

func GetDb() {
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(20)
	db.LogMode(true)
	Engine = db
}

func init() {
	go GetDb()
}
