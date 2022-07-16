package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
	"go-file/middleware"
)

func SetRouter(router *gin.Engine) {
	router.Use(middleware.AllStat())
	setWebRouter(router)
	setApiRouter(router)
	router.NoRoute(controller.Get404Page)
}
