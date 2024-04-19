package model

import (
	"os"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-file/common"
)

var DB *gorm.DB

func createAdminAccount() {
	var user User
	DB.Where(User{Role: common.RoleAdminUser}).Attrs(User{
		Username:    "admin",
		Password:    "123456",
		Role:        common.RoleAdminUser,
		Status:      common.UserStatusEnabled,
		DisplayName: "Administrator",
	}).FirstOrCreate(&user)
}

func CountTable(tableName string) (num int64) {
	DB.Table(tableName).Count(&num)
	return
}

func InitDB() (db *gorm.DB, err error) {
	if os.Getenv("SQL_DSN") != "" {
		// Use MySQL
		//db, err = gorm.Open("mysql", os.Getenv("SQL_DSN"))
		db, err = gorm.Open(mysql.Open(os.Getenv("SQL_DSN")), &gorm.Config{})

	} else {
		// Use SQLite
		//db, err = gorm.Open("sqlite3", common.SQLitePath)
		db, err = gorm.Open(sqlite.Open(common.SQLitePath), &gorm.Config{})
	}
	if err == nil {
		DB = db
		db.AutoMigrate(&File{})
		db.AutoMigrate(&Image{})
		db.AutoMigrate(&User{})
		db.AutoMigrate(&Option{})
		createAdminAccount()
		return DB, err
	} else {
		common.FatalLog("failed to connect to database: " + err.Error())
	}
	return nil, err
}
