package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
	"lan-share/model"
)



func GetIndex(c *gin.Context) {
	query := c.Query("query")

	files, _ := model.Query(query)

	c.HTML(http.StatusOK, "template.gohtml", gin.H{
		"message": "Hi",
		"files":files,
	})
}

func UploadFile(c *gin.Context) {
	description := c.PostForm("description")
	uploader := c.PostForm("uploader")
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Base(file.Filename)
	link := "/upload/"+filename
	if err := c.SaveUploadedFile(file, "./upload/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	fileObj := &model.File{
		Description : description,
		Uploader : uploader,
		Time : currentTime,
		Link : link,
		Filename : filename,
	}
	err = fileObj.Insert()
	if err != nil {
		fmt.Errorf("failed to init database")
	}
	c.Redirect(http.StatusSeeOther, "./")
}

func DeleteFile(c *gin.Context) {
	//id := c.PostForm("id")
	//token := c.PostForm("token")
	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"message": "Token is invalid.",
	})
}