package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"lan-share/model"
	"net/http"
	"path/filepath"
	"time"
)

var uploadPath = "./upload"

type DeleteRequest struct {
	Id    int
	Link  string
	Token string
}

func GetIndex(c *gin.Context) {
	query := c.Query("query")

	files, _ := model.Query(query)

	c.HTML(http.StatusOK, "template.gohtml", gin.H{
		"message": "",
		"files":   files,
	})
}

func UploadFile(c *gin.Context) {
	description := c.PostForm("description")
	if description == "" {
		description = "No description."
	}
	uploader := c.PostForm("uploader")
	if uploader == "" {
		uploader = "Anonymous User"
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Base(file.Filename)
	link := "/upload/" + filename
	if err := c.SaveUploadedFile(file, uploadPath+"/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
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
	if *Token == deleteRequest.Token {
		fileObj := &model.File{
			Id:   deleteRequest.Id,
			Link: deleteRequest.Link,
		}
		err := fileObj.Delete()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "File deleted successfully.",
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "Token is invalid.",
		})
	}
}
