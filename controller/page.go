package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

func GetIndexPage(c *gin.Context) {
	query := c.Query("query")
	isQuery := query != ""
	p, _ := strconv.Atoi(c.Query("p"))
	if p < 0 {
		p = 0
	}
	next := p + 1
	prev := common.IntMax(0, p-1)

	startIdx := p * common.ItemsPerPage

	files, err := model.QueryFiles(query, startIdx)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"message":  err.Error(),
			"option":   common.OptionMap,
			"username": c.GetString("username"),
		})
		return
	}
	if len(files) < common.ItemsPerPage {
		next = 0
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":  "",
		"option":   common.OptionMap,
		"username": c.GetString("username"),
		"files":    files,
		"isQuery":  isQuery,
		"next":     next,
		"prev":     prev,
	})
}

func GetManagePage(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	var uptime = time.Since(common.StartTime)
	session := sessions.Default(c)
	role := session.Get("role")
	c.HTML(http.StatusOK, "manage.html", gin.H{
		"message":                 "",
		"option":                  common.OptionMap,
		"username":                c.GetString("username"),
		"memory":                  fmt.Sprintf("%d MB", m.Sys/1024/1024),
		"uptime":                  common.Seconds2Time(int(uptime.Seconds())),
		"userNum":                 model.CountTable("users"),
		"fileNum":                 model.CountTable("files"),
		"imageNum":                model.CountTable("images"),
		"FileUploadPermission":    common.FileUploadPermission,
		"FileDownloadPermission":  common.FileDownloadPermission,
		"ImageUploadPermission":   common.ImageUploadPermission,
		"ImageDownloadPermission": common.ImageDownloadPermission,
		"isAdmin":                 role == common.RoleAdminUser,
		"StatEnabled":             common.StatEnabled,
	})
}

func GetImagePage(c *gin.Context) {
	c.HTML(http.StatusOK, "image.html", gin.H{
		"message":  "",
		"option":   common.OptionMap,
		"username": c.GetString("username"),
	})
}

func GetLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"message":  "",
		"option":   common.OptionMap,
		"username": c.GetString("username"),
	})
}

func GetHelpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "help.html", gin.H{
		"message":  "",
		"option":   common.OptionMap,
		"username": c.GetString("username"),
	})
}

func Get404Page(c *gin.Context) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"message":  "",
		"option":   common.OptionMap,
		"username": c.GetString("username"),
	})
}
