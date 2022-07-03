package common

import "flag"

var (
	Port  = flag.Int("port", 3000, "specify the server listening port.")
	Token = flag.String("token", "token", "specify the private token.")
	Host  = flag.String("host", "localhost", "the server's ip address or domain")
	Path  = flag.String("path", "", "public a local path")
)

func init() {
	flag.Parse()
}
