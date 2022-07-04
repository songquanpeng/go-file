package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
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

func GetIndexPage(c *gin.Context) {
	query := c.Query("query")

	files, _ := model.QueryFiles(query)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "",
		"files":   files,
	})
}

func GetExplorerPage(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	path, _ = url.PathUnescape(path)

	rootPath := filepath.Join(common.LocalFileRoot, path)
	if !strings.HasPrefix(rootPath, common.LocalFileRoot) {
		// We may being attacked!
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": fmt.Sprintf("You can only access subfolders of the given path."),
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
		for _, f := range files {
			file := model.LocalFile{
				Name:         f.Name(),
				Link:         "explorer?path=" + url.QueryEscape(path+f.Name()),
				Size:         common.Bytes2Size(f.Size()),
				IsFolder:     f.Mode().IsDir(),
				ModifiedTime: f.ModTime().String()[:19],
			}
			if file.IsFolder {
				localFiles = append(localFiles, file)
			} else {
				tempFiles = append(tempFiles, file)
			}

		}
		localFiles = append(localFiles, tempFiles...)

		c.HTML(http.StatusOK, "explorer.html", gin.H{
			"message": "",
			"files":   localFiles,
		})
	} else {
		c.File(filepath.Join(common.LocalFileRoot, path))
	}
}

func GetManagePage(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"message": "",
	})
}

func GetImagePage(c *gin.Context) {
	c.HTML(http.StatusOK, "image.html", gin.H{
		"message": "",
	})
}

func GetLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// TODO: query database to validate username & password, and find his rule
	if username == "admin" && password == *common.Token {
		session := sessions.Default(c)
		session.Set("username", username)
		err := session.Save()
		if err != nil {
			c.HTML(http.StatusForbidden, "login.html", gin.H{
				"message": "Unable to save session, please try again.",
			})
			return
		}
		c.Redirect(http.StatusFound, "/manage")
		return
	} else {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"message": "Wrong user name or password.",
		})
		return
	}
}

func GetVideoPage(c *gin.Context) {
	path := c.DefaultQuery("path", "/")
	path, _ = url.PathUnescape(path)

	rootPath := filepath.Join(common.VideoServePath, path)
	if !strings.HasPrefix(rootPath, common.VideoServePath) {
		// We may being attacked!
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"message": fmt.Sprintf("You can only access subfolders of the given path."),
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
