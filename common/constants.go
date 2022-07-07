package common

import (
	"embed"
	"flag"
	"os"
	"path/filepath"
)

var Version = "v0.4.0"

var (
	Port      = flag.Int("port", 3000, "specify the server listening port.")
	Host      = flag.String("host", "localhost", "the server's ip address or domain")
	Path      = flag.String("path", "", "specify a local path to public")
	VideoPath = flag.String("video", "", "specify a video folder to public")
)

var UploadPath = "upload"
var LocalFileRoot = UploadPath
var ImageUploadPath = "upload/images"
var VideoServePath = "upload"

//go:embed public
var FS embed.FS

var SessionSecret = "I_LOVE_YOU"

func init() {
	flag.Parse()
	if *Path != "" {
		LocalFileRoot = *Path
	}
	if *VideoPath != "" {
		VideoServePath = *VideoPath
	}

	LocalFileRoot, _ = filepath.Abs(LocalFileRoot)
	VideoServePath, _ = filepath.Abs(VideoServePath)
	ImageUploadPath, _ = filepath.Abs(ImageUploadPath)

	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(UploadPath, 0777)
	}
	if _, err := os.Stat(ImageUploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
}
