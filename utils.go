package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
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

func publicLocalPath(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	var files []string
	switch mode := fi.Mode(); {
	case mode.IsDir():
		_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// TODO: unable to public path that start with "."
			// Skip dirs that start with "."
			if info.IsDir() && strings.HasPrefix(path, ".") && !strings.HasPrefix(path, "./") {
				return filepath.SkipDir
			}
			if info.IsDir() {
				return nil
			}
			files = append(files, path)
			return nil
		})
	case mode.IsRegular():
		files = append(files, path)
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	for _, file := range files {
		fileObj := &File{
			Description: file,
			Uploader:    "Local Link",
			Time:        currentTime,
			Link:        "/local/" + file,
			Filename:    filepath.Base(file),
			IsLocalFile: true,
		}
		err = fileObj.Insert()
		if err != nil {
			_ = fmt.Errorf(err.Error())
		}
	}
}

var sizeKB = 1024
var sizeMB = sizeKB * 1024
var sizeGB = sizeMB * 1024
var sizeTB = sizeGB * 1024

func Bytes2Size(num int64) string {
	numStr := ""
	unit := "B"
	if num/int64(sizeTB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeTB))
		unit = "TB"
	} else if num/int64(sizeGB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeGB))
		unit = "GB"
	} else if num/int64(sizeMB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeMB))
		unit = "MB"
	} else if num/int64(sizeKB) > 1 {
		numStr = fmt.Sprintf("%f", float64(num)/float64(sizeKB))
		unit = "KB"
	} else {
		numStr = fmt.Sprintf("%d", num)
	}
	numStr = strings.Split(numStr, ".")[0]
	return numStr + " " + unit
}
