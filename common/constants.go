package common

import "embed"

var UploadPath = "./upload"
var LocalFileRoot = UploadPath

//go:embed public
var FS embed.FS
