package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/controller"
)

func SetRouter(router *gin.Engine) {
	store := cookie.NewStore([]byte(common.SessionSecret))
	router.Use(sessions.Sessions("main", store))

	setWebRouter(router)
	setApiRouter(router)
	router.NoRoute(controller.Get404Page)
}
