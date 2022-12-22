package router

import (
	"embed"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/controller"
	"go-file/middleware"
	"net/http"
)

func setWebRouter(router *gin.Engine, buildFS embed.FS, indexPage []byte) {
	router.Use(middleware.GlobalWebRateLimit())
	fileDownloadRoute := router.Group("/")
	fileDownloadRoute.Use(middleware.FileDownloadPermissionCheck(), middleware.DownloadRateLimit())
	{
		fileDownloadRoute.GET("/upload/:file", controller.DownloadFile)
		fileDownloadRoute.GET("/explorer", controller.GetExplorerPageOrFile)
	}
	router.Use(middleware.Cache())
	router.Use(static.Serve("/", common.EmbedFolder(buildFS, "web/build")))
	router.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexPage)
	})
}
