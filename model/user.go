package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-file/common"
	"os"
	"path/filepath"
)

type User struct {
	Filename string `json:"type" gorm:"type:string"`
	Uploader string `json:"uploader" gorm:"type:string"`
	Time     string `json:"time" gorm:"type:string"`
}

func (user *User) Insert() error {
	var err error
	err = DB.Create(user).Error
	return err
}

func (user *User) Delete() error {
	var err error
	err = DB.Delete(user).Error
	err = os.Remove(filepath.Join(common.ImageUploadPath, user.Filename))
	return err
}
