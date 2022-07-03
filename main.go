package main

import (
	"github.com/gin-gonic/gin"
	"go-file/common"
	"go-file/model"
	"go-file/router"
	"html/template"
	"log"
	"os"
	"strconv"
)

func loadTemplate() *template.Template {
	t := template.Must(template.New("").ParseFS(common.FS, "public/*.html"))
	return t
}

func main() {
	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := model.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	server := gin.Default()
	server.SetHTMLTemplate(loadTemplate())
	router.SetRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*common.Port)
	}
	if *common.Host == "localhost" {
		ip := common.GetIp()
		if ip != "" {
			*common.Host = ip
		}
	}
	serverUrl := "http://" + *common.Host + ":" + realPort + "/"
	common.OpenBrowser(serverUrl)
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}
}
