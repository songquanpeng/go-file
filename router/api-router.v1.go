package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
)

func setApiRouter(router *gin.Engine) {
	router.POST("/upload", controller.UploadFile)
	router.POST("/delete", controller.DeleteFile)
}
