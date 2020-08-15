package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func NewConnection(tbName string) *gorm.DB {
	db, err := gorm.Open("mysql", "root:6e952db19c8436d2@(127.0.0.1:3306)/DMS?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
	}
	return db.Table(tbName)
}
