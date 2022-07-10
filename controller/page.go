package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"net/http"
	"runtime"
	"time"
)

func GetIndexPage(c *gin.Context) {
	query := c.Query("query")
	isQuery := query != ""

	files, _ := model.QueryFiles(query)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "",
		"files":   files,
		"isQuery": isQuery,
	})
}

func GetManagePage(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	var uptime = time.Since(common.StartTime)
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"message":  "",
		"memory":   fmt.Sprintf("%d MB", m.Sys/1024/1024),
		"uptime":   uptime.String(),
		"userNum":  model.CountTable("users"),
		"fileNum":  model.CountTable("files"),
		"imageNum": model.CountTable("images"),
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
