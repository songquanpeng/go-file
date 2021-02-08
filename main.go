package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"html/template"
	"lan-share/model"
	"log"
	"os"
	"strconv"
)

var (
	port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
	Host  = flag.String("host", "localhost", "the server's ip address or domain")
)

var ServerUrl = ""

func init() {
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		_ = os.Mkdir(uploadPath, 0777)
	}
}

func loadTemplate() *template.Template {
	t := template.New("")
	t, err := t.New("template.gohtml").Parse(HTMLTemplate)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return t
}

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
	server.SetHTMLTemplate(loadTemplate())
	SetIndexRouter(server)
	SetApiRouter(server)
	var realPort = os.Getenv("PORT")
	if realPort == "" {
		realPort = strconv.Itoa(*port)
	}
	if *Host == "localhost" {
		ip := getIp()
		if ip != "" {
			*Host = ip
		}
	}
	ServerUrl = "http://" + *Host + ":" + realPort + "/"
	openBrowser(ServerUrl)
	err = server.Run(":" + realPort)
	if err != nil {
		log.Println(err)
	}
}
