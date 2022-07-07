package controller

import (
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"net/http"
)

func GetIndexPage(c *gin.Context) {
	query := c.Query("query")

	files, _ := model.QueryFiles(query)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "",
		"files":   files,
	})
}

func GetManagePage(c *gin.Context) {
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"message": "",
	})
}

func GetImagePage(c *gin.Context) {
	c.HTML(http.StatusOK, "image.html", gin.H{
		"message": "",
	})
}

func GetLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message": "",
	})
}

func GetHelpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", gin.H{
		"message": "",
		"version": common.Version,
	})
}

func Get404Page(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"message": "",
		"version": common.Version,
	})
}
