package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetVideoPage(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	path, _ = url.PathUnescape(path)

	rootPath := filepath.Join(common.VideoServePath, path)
	if !strings.HasPrefix(rootPath, common.VideoServePath) {
		// We may being attacked!
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": fmt.Sprintf("只能访问指定路径下的文件"),
		})
		return
	}
	root, err := os.Stat(rootPath)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	if root.IsDir() {
		var videoPath = ""
		var localFiles []model.LocalFile
		var tempFiles []model.LocalFile
		files, err := ioutil.ReadDir(rootPath)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"message": err.Error(),
			})
			return
		}
		if path != "/" {
			parts := strings.Split(path, "/")
			if len(parts) > 0 {
				parts = parts[:len(parts)-1]
			}
			parentPath := strings.Join(parts, "/")
			parentFile := model.LocalFile{
				Name:         "..",
				Link:         "video?path=" + url.QueryEscape(parentPath),
				Size:         "",
				IsFolder:     true,
				ModifiedTime: "",
			}
			localFiles = append(localFiles, parentFile)
			path = strings.Trim(path, "/") + "/"
		} else {
			path = ""
		}
		for _, f := range files {
			var isFolder = f.Mode().IsDir()
			if !isFolder {
				var ext = filepath.Ext(f.Name())
				if ext != ".mp4" && ext != ".MP4" && ext != ".webm" && ext != ".WEBM" && ext != ".ogg" && ext != ".OGG" {
					continue
				}
			}
			file := model.LocalFile{
				Name:         f.Name(),
				Link:         "video?path=" + url.QueryEscape(path+f.Name()),
				Size:         common.Bytes2Size(f.Size()),
				IsFolder:     isFolder,
				ModifiedTime: f.ModTime().String()[:19],
			}
			if file.IsFolder {
				localFiles = append(localFiles, file)
			} else {
				tempFiles = append(tempFiles, file)
				if videoPath == "" {
					videoPath = file.Link
				}
			}
		}
		localFiles = append(localFiles, tempFiles...)

		c.HTML(http.StatusOK, "video.html", gin.H{
			"message":   "",
			"files":     localFiles,
			"videoPath": videoPath,
		})
	} else {
		c.File(filepath.Join(common.VideoServePath, path))
	}
}
