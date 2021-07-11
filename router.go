package main

import (
	"github.com/gin-gonic/gin"
)

func SetIndexRouter(router *gin.Engine) {
	router.Static("/upload", "./upload")
	router.GET("/local/:path", GetLocalFile)
	router.GET("/public/static/:file", GetStaticFile)
	router.GET("/", GetIndex)
}

func SetApiRouter(router *gin.Engine) {
	router.POST("/upload", UploadFile)
	router.POST("/delete", DeleteFile)
}
