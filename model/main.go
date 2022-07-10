package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func createAdminAccount() {
	var user User
	DB.Where(User{Role: "admin"}).Attrs(User{
		Username:    "admin",
		Password:    "123456",
		Role:        "admin",
		Status:      "active",
		DisplayName: "Administrator",
	}).FirstOrCreate(&user)
}

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./.go-file.db")
	if err == nil {
		DB = db
		db.AutoMigrate(&File{})
		db.AutoMigrate(&Image{})
		db.AutoMigrate(&User{})
		createAdminAccount()
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}
