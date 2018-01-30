package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	var prefix string
	var port string
	flag.StringVar(&prefix, "path", "/", "文件夹前缀")
	flag.StringVar(&port, "port", ":7777", "监听服务和端口")
	flag.Parse()
	fs := new(webdav.Dir)
	h := new(webdav.Handler)
	h.FileSystem = *fs
	h.Prefix = prefix
	h.LockSystem = webdav.NewMemLS()
	//then use the Handler.ServeHTTP Method as the http.HandleFunc
	http.HandleFunc("/", h.ServeHTTP)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Print(err.Error())
	}
}
