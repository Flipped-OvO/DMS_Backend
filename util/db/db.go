package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func NewConnection(tbName string) *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=flipped dbname=DMS sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	return db.Table(tbName)
}
