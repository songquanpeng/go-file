package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-file/common"
	"go-file/model"
	"net/http"
	"path/filepath"
	"time"
)

type ImageDeleteRequest struct {
	Id    string
	Token string
}

func UploadImage(c *gin.Context) {
	uploader := "User" // TODO: check token and find who own it
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": fmt.Sprintf("get form err: %s", err.Error()),
		})
		return
	}
	images := form.File["image"]
	for _, file := range images {
		id := uuid.New().String()
		ext := filepath.Ext(file.Filename)
		if err := c.SaveUploadedFile(file, filepath.Join(common.ImageUploadPath, fmt.Sprintf("%s.%s", id, ext))); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": fmt.Sprintf("upload file err: %s", err.Error()),
			})
			return
		}

		imageObj := &model.Image{
			Id:       id,
			Type:     ext,
			Uploader: uploader,
			Time:     currentTime,
		}
		err = imageObj.Insert()
		if err != nil {
			_ = fmt.Errorf(err.Error())
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
}

func DeleteImage(c *gin.Context) {
	var deleteRequest ImageDeleteRequest
	err := json.NewDecoder(c.Request.Body).Decode(&deleteRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid parameter",
		})
		return
	}
	if *common.Token == deleteRequest.Token {
		imageObj := &model.Image{
			Id: deleteRequest.Id,
		}
		model.DB.Where("id = ?", deleteRequest.Id).First(&imageObj)
		err := imageObj.Delete()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": err.Error(),
			})
		} else {
			message := "Image deleted successfully."
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
