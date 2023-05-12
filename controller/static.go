package controller

import (
	"github.com/gin-gonic/gin"
	"go-file/common"
	"net/http"
)

func GetStaticFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/static/"+path, http.FS(common.FS))
}

func GetLibFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/lib/"+path, http.FS(common.FS))
}

func GetIconFile(c *gin.Context) {
	path := c.Param("file")
	c.FileFromFS("public/icon/"+path, http.FS(common.FS))
}
