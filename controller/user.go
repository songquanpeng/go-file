package controller

import (
	"encoding/json"
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
	session.Set("id", user.Id)
	session.Set("username", username)
	session.Set("role", user.Role)
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

func UpdateSelf(c *gin.Context) {
	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	user.Id = c.GetInt("id")
	role := c.GetString("role")
	if role != "admin" {
		user.Role = ""
		user.Status = ""
	}
	// TODO: check Display Name to avoid XSS attack
	if err := user.Update(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}

// CreateUser Only admin user can call this, so we can trust it
func CreateUser(c *gin.Context) {
	var user model.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	user.DisplayName = user.Username
	// TODO: Check user.Status && user.Role
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}

	if err := user.Insert(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
