package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.HTML(http.StatusForbidden, "login.html", gin.H{
				"message": "请先登录",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("role", session.Get("role"))
		c.Set("id", session.Get("id"))
		c.Next()
	}
}

func ApiAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，请登录后重试",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("role", session.Get("role"))
		c.Set("id", session.Get("id"))
		c.Next()
	}
}

func ApiAdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get("role")
		if role == nil || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，请检查你是否登录或者有相关权限",
			})
			c.Abort()
			return
		}
		c.Set("username", session.Get("username"))
		c.Set("role", role)
		c.Set("id", session.Get("id"))
		c.Next()
	}
}
