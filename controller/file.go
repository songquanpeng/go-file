package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type DeleteRequest struct {
	Id    int
	Link  string
	Token string
}

func UploadFile(c *gin.Context) {
	uploadPath := common.UploadPath
	saveToDatabase := true
	path := c.PostForm("path")
	if path != "" {
		uploadPath = filepath.Join(common.LocalFileRoot, path)
		if !strings.HasPrefix(uploadPath, common.LocalFileRoot) {
			// We may being attacked!
			uploadPath = common.LocalFileRoot
		}
		saveToDatabase = false
	}

	description := c.PostForm("description")
	if description == "" {
		description = "No description."
	}
	uploader := c.PostForm("uploader")
	if uploader == "" {
		uploader = "Anonymous User"
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["file"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		link := "/upload/" + filename
		if err := c.SaveUploadedFile(file, filepath.Join(uploadPath, filename)); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
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
	var deleteRequest DeleteRequest
	err := json.NewDecoder(c.Request.Body).Decode(&deleteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid parameter",
		})
		return
	}
	if *common.Token == deleteRequest.Token {
		fileObj := &model.File{
			Id: deleteRequest.Id,
		}
		model.DB.Where("id = ?", deleteRequest.Id).First(&fileObj)
		err := fileObj.Delete()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": err.Error(),
			})
		} else {
			message := "File deleted successfully."
			if fileObj.IsLocalFile {
				message = "Record deleted successfully."
			}
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": message,
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Token is invalid.",
		})
	}
}
