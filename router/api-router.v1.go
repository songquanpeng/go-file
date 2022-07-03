package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
)

func setApiRouter(router *gin.Engine) {
	router.POST("/file", controller.UploadFile)
	router.DELETE("/file", controller.DeleteFile)

	router.POST("/image", controller.UploadImage)
	router.DELETE("/image", controller.DeleteImage)
}
