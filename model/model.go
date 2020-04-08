package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"strings"
)

type File struct {
	Id              int    `json:"id"`
	Filename        string `json:"filename"`
	Description     string `json:"description"`
	Uploader        string `json:"uploader"`
	Link            string `json:"link"`
	Time            string `json:"time"`
	DownloadCounter int    `json:"download_counter"`
}

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./data.db")
	if err == nil {
		DB = db
		return DB, err
	}
	return nil, err
}

func All() ([]*File, error) {
	var files []*File
	var err error
	err = DB.Find(&files).Error
	return files, err
}

func (file *File) Insert() error {
	var err error
	err = DB.Create(file).Error
	return err
}

func (file *File) Delete() error {
	var err error
	err = DB.Delete(file).Error
	err = os.Remove("." + file.Link)
	return err
}

func Query(query string) ([]*File, error) {
	var files []*File
	var err error
	query = strings.ToLower(query)
	err = DB.Where("filename LIKE ? or description LIKE ? or uploader LIKE ? or time LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&files).Error
	return files, err
}
