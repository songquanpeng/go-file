package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"lan-share/model"
	"os"
	"strconv"
)

var (
	port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()

	db, err := model.InitDB()
	if err != nil {
		fmt.Errorf("failed to init database")
	}
	defer db.Close()
	server := gin.Default()
	server.LoadHTMLGlob("static/template.gohtml")
	SetIndexRouter(server)
	SetApiRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*port)
	}
	_ = server.Run(":" + realPort)
}
