package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
	"go-file/middleware"
)

func setApiRouter(router *gin.Engine) {
	basicAuth := router.Group("/api")
	basicAuth.Use(middleware.ApiAuth())
	{
		basicAuth.POST("/file", controller.UploadFile)
		basicAuth.DELETE("/file", controller.DeleteFile)
		basicAuth.POST("/image", controller.UploadImage)
		basicAuth.DELETE("/image", controller.DeleteImage)
		basicAuth.POST("/user", controller.UpdateUser)
	}
}
