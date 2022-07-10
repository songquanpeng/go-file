package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
	"go-file/middleware"
)

func setApiRouter(router *gin.Engine) {
	fileUploadAuth := router.Group("/api")
	fileUploadAuth.Use(middleware.FileUploadPermissionCheck())
	{
		fileUploadAuth.POST("/file", controller.UploadFile)
	}
	imageUploadAuth := router.Group("/api")
	imageUploadAuth.Use(middleware.ImageUploadPermissionCheck())
	{
		imageUploadAuth.POST("/image", controller.UploadImage)
	}
	basicAuth := router.Group("/api")
	basicAuth.Use(middleware.ApiAuth())
	{
		basicAuth.DELETE("/file", controller.DeleteFile)
		basicAuth.DELETE("/image", controller.DeleteImage)
		basicAuth.PUT("/user", controller.UpdateSelf)
	}
	adminAuth := router.Group("/api")
	adminAuth.Use(middleware.ApiAdminAuth())
	{
		adminAuth.POST("/user", controller.CreateUser)
		adminAuth.PUT("/manage_user", controller.ManageUser)
	}
}
