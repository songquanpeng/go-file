package common

import (
	"flag"
	"path/filepath"
)

var (
	Port      = flag.Int("port", 3000, "specify the server listening port.")
	Token     = flag.String("token", "token", "specify the private token.")
	Host      = flag.String("host", "localhost", "the server's ip address or domain")
	Path      = flag.String("path", "", "specify a local path to public")
	VideoPath = flag.String("video", "", "specify a video folder to public")
)

func init() {
	flag.Parse()
	if *Path != "" {
		LocalFileRoot, _ = filepath.Abs(*Path)
	}
	if *VideoPath != "" {
		VideoServePath, _ = filepath.Abs(*VideoPath)
	}
}
