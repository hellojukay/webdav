package main

import (
	"flag"
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)
var bind *string
var username *string
var password *string
func init(){
	bind = flag.String("bind", ":7777", "监听服务和端口")
	username = flag.String("username","","用户名")
	password = flag.String("password","","密码")
	flag.Parse()
}
func auth(w http.ResponseWriter, req *http.Request){
	u , p , auth := req.BasicAuth()
	if !auth {
		w.WriteHeader(403)
		return
	}
	if !(u == *username && p == *password){
		w.WriteHeader(403)
		return
	}
	var fs webdav.Dir = "/"
	h := new(webdav.Handler)
	h.FileSystem = fs
	h.LockSystem = webdav.NewMemLS()
	h.ServeHTTP(w,req)
}
func main() {
	http.HandleFunc("/",auth)
	//then use the Handler.ServeHTTP Method as the http.HandleFunc
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Print(err.Error())
	}
}
