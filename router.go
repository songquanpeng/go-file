package main

import (
	"github.com/gin-gonic/gin"
)

func SetIndexRouter(router *gin.Engine) {
	router.Static("/upload", "./upload")
	router.GET("/explorer", GetExplorerIndex)
	router.GET("/public/static/:file", GetStaticFile)
	router.GET("/public/lib/:file", GetLibFile)
	router.GET("/", GetIndex)
}

func SetApiRouter(router *gin.Engine) {
	router.POST("/upload", UploadFile)
	router.POST("/delete", DeleteFile)
}
