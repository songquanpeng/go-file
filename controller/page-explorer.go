package controller

import (
	"context"
	"encoding/json"
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
	"time"
)

func GetExplorerPageOrFile(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	path, _ = url.PathUnescape(path)

	fullPath := filepath.Join(common.ExplorerRootPath, path)
	if !strings.HasPrefix(fullPath, common.ExplorerRootPath) {
		// We may being attacked!
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message":  fmt.Sprintf("只能访问指定文件夹的子目录"),
			"option":   common.OptionMap,
			"username": c.GetString("username"),
		})
		return
	}
	root, err := os.Stat(fullPath)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message":  "处理路径时发生了错误，请确认路径正确",
			"option":   common.OptionMap,
			"username": c.GetString("username"),
		})
		return
	}
	if root.IsDir() {
		localFilesPtr, readmeFileLink, err := getData(path, fullPath)
		if err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{
				"message":  err.Error(),
				"option":   common.OptionMap,
				"username": c.GetString("username"),
			})
			return
		}

		c.HTML(http.StatusOK, "explorer.html", gin.H{
			"message":        "",
			"option":         common.OptionMap,
			"username":       c.GetString("username"),
			"files":          localFilesPtr,
			"readmeFileLink": readmeFileLink,
		})
	} else {
		c.File(filepath.Join(common.ExplorerRootPath, path))
	}
}

func getDataFromFS(path string, fullPath string) (localFilesPtr *[]model.LocalFile, readmeFileLink string, err error) {
	var localFiles []model.LocalFile
	var tempFiles []model.LocalFile
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return
	}
	if path != "/" {
		parts := strings.Split(path, "/")
		// Add the special item: ".." which means parent dir
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
	localFilesPtr = &localFiles
	return
}

func getData(path string, fullPath string) (localFilesPtr *[]model.LocalFile, readmeFileLink string, err error) {
	if !common.ExplorerCacheEnabled {
		return getDataFromFS(path, fullPath)
	} else {
		ctx := context.Background()
		rdb := common.RDB
		key := "cacheExplorer:" + fullPath
		n, _ := rdb.Exists(ctx, key).Result()
		if n <= 0 {
			// Cache doesn't exist
			localFilesPtr, readmeFileLink, err = getDataFromFS(path, fullPath)
			if err != nil {
				return
			}
			// Start a coroutine to update cache
			go func() {
				var values []string
				for _, f := range *localFilesPtr {
					s, err := json.Marshal(f)
					if err != nil {
						return
					}
					values = append(values, string(s))
				}
				rdb.RPush(ctx, key, values)
				rdb.Expire(ctx, key, time.Duration(common.ExplorerCacheTimeout)*time.Second)
			}()
		} else {
			// Cache existed, use cached data
			var localFiles []model.LocalFile
			file := model.LocalFile{}
			for _, s := range rdb.LRange(ctx, key, 0, -1).Val() {
				err = json.Unmarshal([]byte(s), &file)
				if err != nil {
					return
				}
				if file.Name == "README.md" {
					readmeFileLink = file.Link
				}
				localFiles = append(localFiles, file)
			}
			localFilesPtr = &localFiles
		}
	}
	return
}
