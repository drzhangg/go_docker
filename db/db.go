package db

import (
	"demo/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var Engine *gorm.DB

var dataSource = "%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

func GetDb() {
	dataSource = fmt.Sprintf(dataSource, config.DBConfig().User, config.DBConfig().Password, config.DBConfig().Host, config.DBConfig().Port, config.DBConfig().Database)
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
