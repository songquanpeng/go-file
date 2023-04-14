package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileDeleteRequest struct {
	Id   int
	Link string
	//Token string
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
	if uploader == "" {
		uploader = "匿名用户"
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
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
		link := filename
		savePath := filepath.Join(uploadPath, filename)
		if _, err := os.Stat(savePath); err == nil {
			// File already existed.
			t := time.Now()
			timestamp := t.Format("_2006-01-02_15-04-05")
			ext := filepath.Ext(filename)
			if ext == "" {
				link += timestamp
			} else {
				link = filename[:len(filename)-len(ext)] + timestamp + ext
			}
			savePath = filepath.Join(uploadPath, link)
		}
		if createTextFile {
			// Create a new text file and then write the description to it.
			filename = "文本分享"
			f, err := os.Create(savePath)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to create file: %s", err.Error()))
				return
			}
			_, err = f.WriteString(description)
			if err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to write text to file: %s", err.Error()))
				return
			}
			descriptionRune := []rune(description)
			if len(descriptionRune) > common.AbstractTextLength {
				description = fmt.Sprintf("内容摘要：%s...", string(descriptionRune[:common.AbstractTextLength]))
			}
		} else {
			if err := c.SaveUploadedFile(file, savePath); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to save uploaded file: %s", err.Error()))
				return
			}
		}
		if saveToDatabase {
			fileObj := &model.File{
				Description: description,
				Uploader:    uploader,
				Time:        currentTime,
				Link:        link,
				Filename:    filename,
			}
			err = fileObj.Insert()
			if err != nil {
				_ = fmt.Errorf(err.Error())
			}
		}
	}
	c.Redirect(http.StatusSeeOther, "./")
}

func DeleteFile(c *gin.Context) {
	var deleteRequest FileDeleteRequest
	err := json.NewDecoder(c.Request.Body).Decode(&deleteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}

	fileObj := &model.File{
		Id: deleteRequest.Id,
	}
	model.DB.Where("id = ?", deleteRequest.Id).First(&fileObj)
	err = fileObj.Delete()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": err.Error(),
		})
	} else {
		message := "文件删除成功"
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": message,
		})
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
	if strings.HasSuffix(fullPath, ".txt") && common.IsMobileUserAgent(c.Request.UserAgent()) {
		content, err := os.ReadFile(fullPath)
		if err != nil {
			c.Status(404)
			return
		}
		c.HTML(http.StatusOK, "text-copy.html", gin.H{
			"content": string(content),
		})
	} else {
		c.File(fullPath)
	}
	// Update download counter
	go func() {
		model.UpdateDownloadCounter(path)
	}()
}
