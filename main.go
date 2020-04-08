package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	server := gin.Default()
	server.LoadHTMLGlob("static/template.gohtml")
	SetIndexRouter(server)
	SetApiRouter(server)
	_ = server.Run(":3000")
}