package main

import (
	"log"
	"net/http"
	router "project1/routers"
	"project1/utls"

	"github.com/urfave/negroni"
)

func main() {
	
	r := router.InitRoutes()
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(corsMiddleware))
	n.With()
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

func corsMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	// Call the next handler
	next(w, r)
}