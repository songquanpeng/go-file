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
	t := time.Now()
	subfolder := t.Format("2006-01")
	err = common.MakeDirIfNotExist(filepath.Join(uploadPath, subfolder))
	if err != nil {
		common.SysError("failed to create folder: " + err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	for _, file := range files {
		// In case someone wants to upload to other folders.
		filename := filepath.Base(file.Filename)
		link := fmt.Sprintf("%s/%s", subfolder, filename)
		savePath := filepath.Join(uploadPath, subfolder, filename)
		if _, err := os.Stat(savePath); err == nil {
			// File already existed.
			timestamp := t.Format("_2006-01-02_15-04-05")
			ext := filepath.Ext(filename)
			if ext == "" {
				link += timestamp
			} else {
				link = subfolder + "/" + filename[:len(filename)-len(ext)] + timestamp + ext
			}
			savePath = filepath.Join(uploadPath, link)
		}
		if createTextFile {
			// Create a new text file and then write the description to it.
			filename = "文本分享"
			f, err := os.Create(savePath)
			if err != nil {
				message := "failed to create file: " + err.Error()
				common.SysError(message)
				c.String(http.StatusInternalServerError, message)
				return
			}
			_, err = f.WriteString(description)
			if err != nil {
				message := "failed to write text to file: " + err.Error()
				common.SysError(message)
				c.String(http.StatusInternalServerError, message)
				return
			}
			descriptionRune := []rune(description)
			if len(descriptionRune) > common.AbstractTextLength {
				description = fmt.Sprintf("内容摘要：%s...", string(descriptionRune[:common.AbstractTextLength]))
			}
		} else {
			if err := c.SaveUploadedFile(file, savePath); err != nil {
				message := "failed to save uploaded file: " + err.Error()
				common.SysError(message)
				c.String(http.StatusInternalServerError, message)
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
				common.SysError("failed to insert file to database: " + err.Error())
				continue
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
	path := c.Param("filepath")
	subfolder, filename := filepath.Split(path)
	link := filename // Keep compatibility with old version
	if subfolder != "/" {
		link = fmt.Sprintf("%s%s", subfolder, filename)
		link = strings.TrimPrefix(link, "/")
	}
	fullPath := filepath.Join(common.UploadPath, subfolder, filename)
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
		model.UpdateDownloadCounter(link)
	}()
}
