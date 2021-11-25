package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

	files, _ := Query(query)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "",
		"files":   files,
	})
}

func GetLocalFile(c *gin.Context) {
	fileObj := &File{}
	path := c.Param("path")
	DB.Where("link = ?", "/local/"+path).First(&fileObj)
	if fileObj.IsLocalFile {
		c.File(path)
	} else {
		c.AbortWithStatus(404)
	}
}

func GetStaticFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/static/"+path, http.FS(fs))
}

func GetLibFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/lib/"+path, http.FS(fs))
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
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["file"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		link := "/upload/" + filename
		if err := c.SaveUploadedFile(file, uploadPath+"/"+filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		fileObj := &File{
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
		fileObj := &File{
			Id: deleteRequest.Id,
		}
		DB.Where("id = ?", deleteRequest.Id).First(&fileObj)
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
