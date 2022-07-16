package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-file/common"
	"os"
	"strings"
)

type File struct {
	Id              int    `json:"id"`
	Filename        string `json:"filename" gorm:"type:string"`
	Description     string `json:"description" gorm:"type:string"`
	Uploader        string `json:"uploader" gorm:"type:string"`
	Link            string `json:"link" gorm:"type:string unique"`
	Time            string `json:"time" gorm:"type:string"`
	DownloadCounter int    `json:"download_counter" gorm:"type:int"`
	IsLocalFile     bool   `json:"is_local_file" gorm:"type:bool"`
}

type LocalFile struct {
	Name         string
	Link         string
	Size         string
	IsFolder     bool
	ModifiedTime string
}

func AllFiles() ([]*File, error) {
	var files []*File
	var err error
	err = DB.Find(&files).Error
	return files, err
}

func QueryFiles(query string, startIdx int) ([]*File, error) {
	var files []*File
	var err error
	query = strings.ToLower(query)
	err = DB.Limit(common.ItemsPerPage).Offset(startIdx).Where("filename LIKE ? or description LIKE ? or uploader LIKE ? or time LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Order("id desc").Find(&files).Error
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
	if !file.IsLocalFile {
		err = os.Remove("." + file.Link)
	}
	return err
}
