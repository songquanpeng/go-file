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
			"message":  fmt.Sprintf("只能访问指定路径下的文件"),
			"option":   common.OptionMap,
			"username": c.GetString("username"),
		})
		return
	}
	root, err := os.Stat(rootPath)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message":  err.Error(),
			"option":   common.OptionMap,
			"username": c.GetString("username"),
		})
		return
	}
	if root.IsDir() {
		var videoPath = ""
		var videoName = "请选择视频进行播放"
		var localFiles []model.LocalFile
		var tempFiles []model.LocalFile
		files, err := ioutil.ReadDir(rootPath)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"message":  err.Error(),
				"option":   common.OptionMap,
				"username": c.GetString("username"),
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
				Name:         "上级目录",
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
			filename := f.Name()
			var isFolder = f.Mode().IsDir()
			if !isFolder {
				var ext = filepath.Ext(f.Name())
				if ext != ".mp4" && ext != ".MP4" && ext != ".webm" && ext != ".WEBM" &&
					ext != ".ogg" && ext != ".OGG" && ext != ".mkv" && ext != ".MKV" {
					continue
				}
				filename = strings.TrimSuffix(filename, ext)
				filename = strings.ReplaceAll(filename, ".", " ")
			}
			file := model.LocalFile{
				Name:         filename,
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
					videoName = filename
				}
			}
		}
		localFiles = append(localFiles, tempFiles...)

		c.HTML(http.StatusOK, "video.html", gin.H{
			"message":   "",
			"option":    common.OptionMap,
			"username":  c.GetString("username"),
			"files":     localFiles,
			"videoPath": videoPath,
			"videoName": videoName,
		})
	} else {
		c.File(filepath.Join(common.VideoServePath, path))
	}
}
