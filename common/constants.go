package common

import (
	"embed"
	"os"
)

var UploadPath = "upload"
var LocalFileRoot = UploadPath
var ImageUploadPath = "upload/image"

//go:embed public
var FS embed.FS

func init() {
	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(UploadPath, 0777)
	}
	if _, err := os.Stat(ImageUploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
}
