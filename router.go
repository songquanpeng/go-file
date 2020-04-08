package main

import (
	"github.com/gin-gonic/gin"
)

func SetIndexRouter(router *gin.Engine) {
	router.Static("/static", "./static")
	router.Static("/upload", "./upload")
	router.GET("/", GetIndex)
}

func SetApiRouter(router *gin.Engine) {
	router.POST("/upload", UploadFile)
	router.DELETE("/", DeleteFile)
}