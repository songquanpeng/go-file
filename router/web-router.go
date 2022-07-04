package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/controller"
	"go-file/middleware"
)

func setWebRouter(router *gin.Engine) {
	store := cookie.NewStore([]byte(common.SessionSecret))
	router.Use(sessions.Sessions("main", store))

	// Always available
	router.GET("/", controller.GetIndexPage)
	router.GET("/public/static/:file", controller.GetStaticFile)
	router.GET("/public/lib/:file", controller.GetLibFile)
	router.GET("/login", controller.GetLoginPage)
	router.POST("/login", controller.Login)

	// Download files
	router.Static("/upload", common.UploadPath)
	router.Static("/image", common.ImageUploadPath)

	router.GET("/explorer", controller.GetExplorerPage)
	router.GET("/image", controller.GetImagePage)

	router.GET("/video", controller.GetVideoPage)

	basicAuth := router.Group("/")
	basicAuth.Use(middleware.WebAuth())
	{
		basicAuth.GET("/manage", controller.GetManagePage)
	}
}
