package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-file/common"
	"os"
	"path/filepath"
)

type Image struct {
	Filename string `json:"type"`
	Uploader string `json:"uploader"`
	Time     string `json:"time"`
}

func AllImage() ([]*Image, error) {
	var images []*Image
	var err error
	err = DB.Find(&images).Error
	return images, err
}

func (image *Image) Insert() error {
	var err error
	err = DB.Create(image).Error
	return err
}

func (image *Image) Delete() error {
	var err error
	err = DB.Delete(image).Error
	err = os.Remove(filepath.Join(common.ImageUploadPath, image.Filename))
	return err
}
