package common

import (
	"embed"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var StartTime = time.Now()
var Version = "v0.0.0"
var OptionMap map[string]string

var ItemsPerPage = 10
var AbstractTextLength = 40

var ExplorerCacheEnabled = false // After my test, enable this will make the server slower...
var ExplorerCacheTimeout = 600   // Second

var StatEnabled = true
var StatCacheTimeout = 24 // Hour
var StatReqTimeout = 30   // Day
var StatIPNum = 20
var StatURLNum = 20

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
	Port         = flag.Int("port", 3000, "Specify the server listening port.")
	Host         = flag.String("host", "", "The server's IP address or domain.")
	Path         = flag.String("path", "", "Specify a local path to public.")
	VideoPath    = flag.String("video", "", "Specify a folder containing videos to be made public.")
	NoBrowser    = flag.Bool("no-browser", false, "Do not open browser automatically.")
	PrintVersion = flag.Bool("version", false, "Print version information.")
	EnableP2P    = flag.Bool("enable-p2p", false, "Enable peer-to-peer relay or not.")
	P2PPort      = flag.Int("p2p-port", 9377, "Specify the P2P listening port.")
	LogDir       = flag.String("log-dir", "", "Specify the directory for log files.")
	PrintHelp    = flag.Bool("help", false, "Print usage information.")
)

// UploadPath Maybe override by ENV_VAR
var UploadPath = "upload"
var ExplorerRootPath = UploadPath
var ImageUploadPath = "upload/images"
var VideoServePath = "upload"

//go:embed public
var FS embed.FS

var SessionSecret = uuid.New().String()

var SQLitePath = "go-file.db"

func printHelp() {
	fmt.Println(fmt.Sprintf("Go File %s - A simple file sharing tool.", Version))
	fmt.Println("Copyright (C) 2023 JustSong. All rights reserved.")
	fmt.Println("GitHub: https://github.com/songquanpeng/go-file")
	fmt.Println("Usage: go-file [options]")
	fmt.Println("Options:")
	flag.CommandLine.VisitAll(func(f *flag.Flag) {
		name := fmt.Sprintf("-%s", f.Name)
		usage := strings.Replace(f.Usage, "\n", "\n    ", -1)
		fmt.Printf("        -%-14s%s\n", name, usage)
	})
	os.Exit(0)
}

func init() {
	flag.Parse()

	if *PrintHelp {
		printHelp()
	}

	if *PrintVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	if os.Getenv("SESSION_SECRET") != "" {
		SessionSecret = os.Getenv("SESSION_SECRET")
	}
	if os.Getenv("SQLITE_PATH") != "" {
		SQLitePath = os.Getenv("SQLITE_PATH")
	}
	if os.Getenv("UPLOAD_PATH") != "" {
		UploadPath = os.Getenv("UPLOAD_PATH")
		ExplorerRootPath = UploadPath
		ImageUploadPath = path.Join(UploadPath, "images")
		VideoServePath = UploadPath
	}
	if *Path != "" {
		ExplorerRootPath = *Path
	}
	if *VideoPath != "" {
		VideoServePath = *VideoPath
	}

	ExplorerRootPath, _ = filepath.Abs(ExplorerRootPath)
	VideoServePath, _ = filepath.Abs(VideoServePath)
	ImageUploadPath, _ = filepath.Abs(ImageUploadPath)

	if _, err := os.Stat(UploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(UploadPath, 0777)
	}
	if _, err := os.Stat(ImageUploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(ImageUploadPath, 0777)
	}
}
