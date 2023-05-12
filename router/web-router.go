package router

import (
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/controller"
	"go-file/middleware"
)

func setWebRouter(router *gin.Engine) {
	router.Use(middleware.GlobalWebRateLimit())
	// Always available
	// All page must have username in context
	router.GET("/", middleware.ExtractUserInfo(), controller.GetIndexPage)
	router.GET("/public/static/:file", controller.GetStaticFile)
	router.GET("/public/lib/:file", controller.GetLibFile)
	router.GET("/public/icon/:file", controller.GetIconFile)
	router.GET("/login", middleware.ExtractUserInfo(), controller.GetLoginPage)
	router.POST("/login", middleware.CriticalRateLimit(), controller.Login)
	router.GET("/logout", controller.Logout)
	router.GET("/help", middleware.ExtractUserInfo(), controller.GetHelpPage)

	// Download files
	fileDownloadAuth := router.Group("/")
	fileDownloadAuth.Use(middleware.DownloadRateLimit(), middleware.FileDownloadPermissionCheck())
	{
		fileDownloadAuth.GET("/upload/*filepath", controller.DownloadFile)
		fileDownloadAuth.GET("/explorer", middleware.ExtractUserInfo(), controller.GetExplorerPageOrFile)
	}

	imageDownloadAuth := router.Group("/")
	imageDownloadAuth.Use(middleware.DownloadRateLimit(), middleware.ImageDownloadPermissionCheck())
	{
		imageDownloadAuth.Static("/image", common.ImageUploadPath)
	}

	router.GET("/image", middleware.ExtractUserInfo(), controller.GetImagePage)

	router.GET("/video", middleware.ExtractUserInfo(), controller.GetVideoPage)

	basicAuth := router.Group("/")
	basicAuth.Use(middleware.WebAuth()) // WebAuth already has username in context
	{
		basicAuth.GET("/manage", controller.GetManagePage)
	}
}
