package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
)

func setWebRouter(router *gin.Engine) {
	router.Static("/upload", "./upload")
	router.GET("/explorer", controller.GetExplorerPage)
	router.GET("/manage", controller.GetManagePage)
	router.GET("/image", controller.GetImagePage)
	router.GET("/public/static/:file", controller.GetStaticFile)
	router.GET("/public/lib/:file", controller.GetLibFile)
	router.GET("/", controller.GetIndexPage)
}
