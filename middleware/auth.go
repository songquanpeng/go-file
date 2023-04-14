package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"net/http"
)

func WebAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.HTML(http.StatusForbidden, "login.html", gin.H{
				"message": "未登录或登录已过期",
				"option":  common.OptionMap,
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

func ExtractUserInfo() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			username = ""
		}
		c.Set("username", username)
		c.Next()
	}
}

func ApiAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		role := session.Get("role")
		id := session.Get("id")
		if username == nil {
			// Check token
			token := c.Request.Header.Get("Authorization")
			user := model.ValidateUserToken(token)
			if user != nil && user.Username != "" {
				// Token is valid
				username = user.Username
				role = user.Role
				id = user.Id
			} else {
				c.JSON(http.StatusForbidden, gin.H{
					"success": false,
					"message": "无权进行此操作，未登录或 token 无效",
				})
				c.Abort()
				return
			}
		}
		c.Set("username", username)
		c.Set("role", role)
		c.Set("id", id)
		c.Next()
	}
}

func ApiAdminAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		role := session.Get("role")
		id := session.Get("id")
		if username == nil {
			// Check token
			token := c.Request.Header.Get("Authorization")
			user := model.ValidateUserToken(token)
			if user != nil && user.Username != "" {
				// Token is valid
				username = user.Username
				role = user.Role
				id = user.Id
			}
		}
		if role != common.RoleAdminUser {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，未登录或 token 无效，或没有权限",
			})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Set("role", role)
		c.Set("id", id)
		c.Next()
	}
}

func NoTokenAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authByToken := c.GetString("authByToken")
		if authByToken == "true" {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "该接口不能使用 token 进行验证",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
