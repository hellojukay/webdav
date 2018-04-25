package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)
var bind *string
var h *webdav.Handler

func main() {
	bind = flag.String("bind", ":7777", "监听服务和端口")
	
	flag.Parse()
	var fs webdav.Dir = "/"
	h := new(webdav.Handler)
	h.FileSystem = fs
	h.LockSystem = webdav.NewMemLS()
	http.HandleFunc("/",h.ServeHTTP)
	//then use the Handler.ServeHTTP Method as the http.HandleFunc
	err := http.ListenAndServe(*bind, nil)
	if err != nil {
		log.Print(err.Error())
	}
}
