package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"lan-share/model"
	"log"
	"os"
	"strconv"
)

var (
	port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
)

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	flag.Parse()

	db, err := model.InitDB()
	if err != nil {
		log.Fatal(err)
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
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}
}
