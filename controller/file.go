package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetAllFiles(c *gin.Context) {
	p, _ := strconv.Atoi(c.Query("p"))
	if p < 0 {
		p = 0
	}
	files, err := model.GetAllFiles(p*common.ItemsPerPage, common.ItemsPerPage)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    files,
	})
	return
}

func SearchFiles(c *gin.Context) {
	keyword := c.Query("keyword")
	files, err := model.SearchFiles(keyword)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    files,
	})
	return
}

func UploadFile(c *gin.Context) {
	uploadPath := common.UploadPath
	saveToDatabase := true
	path := c.PostForm("path")
	if path != "" { // Upload to explorer's path
		uploadPath = filepath.Join(common.ExplorerRootPath, path)
		if !strings.HasPrefix(uploadPath, common.ExplorerRootPath) {
			// In this case the given path is not valid, so we reset it to ExplorerRootPath.
			uploadPath = common.ExplorerRootPath
		}
		saveToDatabase = false

		// Start a go routine to delete explorer' cache
		if common.ExplorerCacheEnabled {
			go func() {
				ctx := context.Background()
				rdb := common.RDB
				key := "cacheExplorer:" + uploadPath
				rdb.Del(ctx, key)
			}()
		}
	}

	description := c.PostForm("description")
	uploader := c.GetString("username")
	uploaderId := c.GetInt("id")
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	files := form.File["file"]
	createTextFile := false
	if files == nil && description != "" {
		createTextFile = true
		file := &multipart.FileHeader{
			Filename: "text.txt",
			Header:   nil,
			Size:     0,
		}
		files = append(files, file)
	}
	for _, file := range files {
		// In case someone wants to upload to other folders.
		filename := filepath.Base(file.Filename)
		ext := filepath.Ext(filename)
		link := common.GetUUID() + ext
		savePath := filepath.Join(uploadPath, link) // both parts are checked, so this path should be safe to use
		if createTextFile {
			// Create a new text file and then write the description to it.
			filename = "文本分享"
			f, err := os.Create(savePath)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": fmt.Sprintf("failed to create file: %s", err.Error()),
				})
				return
			}
			_, err = f.WriteString(description)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": fmt.Sprintf("failed to write text to file: %s", err.Error()),
				})
				return
			}
			descriptionRune := []rune(description)
			if len(descriptionRune) > common.AbstractTextLength {
				description = fmt.Sprintf("内容摘要：%s...", string(descriptionRune[:common.AbstractTextLength]))
			}
		} else {
			if err := c.SaveUploadedFile(file, savePath); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": err.Error(),
				})
				return
			}
		}
		// save to database
		if saveToDatabase {
			fileObj := &model.File{
				Description: description,
				Uploader:    uploader,
				UploadTime:  currentTime,
				UploaderId:  uploaderId,
				Link:        link,
				Filename:    filename,
			}
			err = fileObj.Insert()
			if err != nil {
				_ = fmt.Errorf(err.Error())
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}

func DeleteFile(c *gin.Context) {
	fileIdStr := c.Param("id")
	fileId, err := strconv.Atoi(fileIdStr)
	if err != nil || fileId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}

	fileObj := &model.File{
		Id: fileId,
	}
	model.DB.Where("id = ?", fileId).First(&fileObj)
	if fileObj.Link == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "文件不存在！",
		})
		return
	}
	err = fileObj.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": err.Error(),
		})
		return
	} else {
		message := "文件删除成功"
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": message,
		})
		return
	}
}

func DownloadFile(c *gin.Context) {
	path := c.Param("file")
	fullPath := filepath.Join(common.UploadPath, path)
	if !strings.HasPrefix(fullPath, common.UploadPath) {
		// We may being attacked!
		c.Status(403)
		return
	}
	c.File(fullPath)
	// Update download counter
	go func() {
		model.UpdateDownloadCounter(path)
	}()
}
