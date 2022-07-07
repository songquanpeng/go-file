package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-file/model"
	"net/http"
	"strings"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	user.ValidateAndFill()
	if user.Status != "active" {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"message": "用户名或密码错误，或者该用户已被封禁",
		})
		return
	}

	session := sessions.Default(c)
	session.Set("username", username)
	session.Set("rule", user.Role)
	err := session.Save()
	if err != nil {
		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"message": "无法保存会话信息，请重试",
		})
		return
	}
	redirectUrl := c.Request.Referer()
	if strings.HasSuffix(redirectUrl, "/login") {
		redirectUrl = "/"
	}
	c.Redirect(http.StatusFound, redirectUrl)
	return
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
