package main

import (
	"log"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	//case "linux":
	//	err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	if err != nil {
		log.Println(err)
	}
}

func getIp() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return ip
	}

	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				if strings.HasPrefix(ip, "192.168") && !strings.HasSuffix(ip, ".1") {
					return
				}
				ip = ""
			}
		}
	}
	return
}
