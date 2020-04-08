package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type File struct {
	Id string
	Filename string
	Description string
	Uploader string
	Link string
	Time string
	DownloadCounter int
}

func GetIndex(c *gin.Context) {
	//query := c.Query("query")
	file := &File{
		Id:              "1",
		Filename:        "README.md",
		Description:     "description",
		Uploader:        "uploader",
		Link:            "/upload/README.md",
		Time:            "2020-04-05",
		DownloadCounter: 2,
	}
	c.HTML(http.StatusOK, "template.gohtml", gin.H{
		"message":"",
		"files":[2]*File{file, file},
	})
}

func UploadFile(c *gin.Context) {
	//description := c.PostForm("description")
	//uploader := c.PostForm("uploader")
	//time := time.Now().Format("2006-01-02 15:04:05")
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Base(file.Filename)
	//link := "/upload/"+filename
	if err := c.SaveUploadedFile(file, "./upload/"+filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
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