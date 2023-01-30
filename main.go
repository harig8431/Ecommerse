package main

import (
	"log"
	"net/http"
	router "project/routers"
	"project/utls"

	"github.com/urfave/negroni"
)

func main() {
	
	r := router.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(r)

	serverip,serverport:=Readserversetting()

	server:=&http.Server{
		Addr: serverip +":"+serverport,
		Handler: n,
	}
log.Println("Listening on",server.Addr)
err:=server.ListenAndServe()
if err!=nil{
	log.Println(err)
}
}

func Readserversetting() (string, string) {

	ConfigFile:=utls.LoadConfiguration()

	return ConfigFile.ServerIp,ConfigFile.ServerPort
}