package router

import "github.com/gin-gonic/gin"

func SetRouter(router *gin.Engine) {
	setWebRouter(router)
	setApiRouter(router)
}
