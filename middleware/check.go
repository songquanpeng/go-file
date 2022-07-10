package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"net/http"
)

func permissionCheckHelper(c *gin.Context, requiredPermission int) {
	c.Set("username", "访客")
	c.Set("role", common.RoleGuestUser)
	c.Set("id", 0)
	if requiredPermission == common.RoleGuestUser {
		c.Next()
		return
	}
	session := sessions.Default(c)
	role := session.Get("role")
	if role == nil || role.(int) < requiredPermission {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "无权进行此操作，请检查你是否登录或者有相关权限",
		})
		c.Abort()
		return
	}
	c.Next()
}

func ImageDownloadPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		permissionCheckHelper(c, common.ImageDownloadPermission)
	}
}

func ImageUploadPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		permissionCheckHelper(c, common.ImageUploadPermission)
	}
}

func FileDownloadPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		permissionCheckHelper(c, common.FileDownloadPermission)
	}
}

func FileUploadPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		permissionCheckHelper(c, common.FileUploadPermission)
	}
}
