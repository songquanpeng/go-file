package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lan-share/model"
)

func main()  {
	db, err := model.InitDB()
	if err != nil {
		fmt.Errorf("failed to init database")
	}
	defer db.Close()
	server := gin.Default()
	server.LoadHTMLGlob("static/template.gohtml")
	SetIndexRouter(server)
	SetApiRouter(server)
	_ = server.Run(":3000")
}