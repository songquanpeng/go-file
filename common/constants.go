package common

import (
	"embed"
	"flag"
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"time"
)

var StartTime = time.Now()
var Version = "v0.4.0"
var OptionMap map[string]interface{}

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
)

var (
	FileUploadPermission    = RoleGuestUser
	FileDownloadPermission  = RoleGuestUser
	ImageUploadPermission   = RoleGuestUser
	ImageDownloadPermission = RoleGuestUser
)

var (
	GlobalApiRateLimit = 20
	GlobalWebRateLimit = 60
	DownloadRateLimit  = 10
	CriticalRateLimit  = 3
)

const (
	UserStatusEnabled  = 1
	UserStatusDisabled = 2 // don't use 0
)

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

var SessionSecret = uuid.New().String()

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

	// TODO Initialize OptionMap
	//
	//OptionMap[""] = ""
}
