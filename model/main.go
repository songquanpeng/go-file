package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./.go-file.db")
	if err == nil {
		DB = db
		db.AutoMigrate(&File{})
		db.AutoMigrate(&Image{})
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}
