package router

import (
	"github.com/gin-gonic/gin"
	"go-file/controller"
)

func SetRouter(router *gin.Engine) {
	setWebRouter(router)
	setApiRouter(router)
	router.NoRoute(controller.Get404Page)
}
