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

func GetExplorerPageOrFile(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	path, _ = url.PathUnescape(path)

	rootPath := filepath.Join(common.LocalFileRoot, path)
	if !strings.HasPrefix(rootPath, common.LocalFileRoot) {
		// We may being attacked!
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": fmt.Sprintf("只能访问指定文件夹的子目录"),
		})
		return
	}
	root, err := os.Stat(rootPath)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": "处理路径时发生了错误，请确认路径正确",
		})
		return
	}
	if root.IsDir() {
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
				Link:         "explorer?path=" + url.QueryEscape(parentPath),
				Size:         "",
				IsFolder:     true,
				ModifiedTime: "",
			}
			localFiles = append(localFiles, parentFile)
			path = strings.Trim(path, "/") + "/"
		} else {
			path = ""
		}
		readmeFileLink := ""
		for _, f := range files {
			link := "explorer?path=" + url.QueryEscape(path+f.Name())
			file := model.LocalFile{
				Name:         f.Name(),
				Link:         link,
				Size:         common.Bytes2Size(f.Size()),
				IsFolder:     f.Mode().IsDir(),
				ModifiedTime: f.ModTime().String()[:19],
			}
			if file.IsFolder {
				localFiles = append(localFiles, file)
			} else {
				tempFiles = append(tempFiles, file)
			}
			if f.Name() == "README.md" {
				readmeFileLink = link
			}
		}
		localFiles = append(localFiles, tempFiles...)

		c.HTML(http.StatusOK, "explorer.html", gin.H{
			"message":        "",
			"files":          localFiles,
			"readmeFileLink": readmeFileLink,
		})
	} else {
		c.File(filepath.Join(common.LocalFileRoot, path))
	}
}
