package main

import (
	"log"
	"net/http"
	"os"
	"proxy/handler"
	"proxy/server"
)

var (
	port = "8083"
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
}

func main() {
	server := server.Server{
		Router: http.NewServeMux(),
	}
	h := handler.Handler{}
	server.InitRoute(&h)
	log.Println("started proxy microservice ...")
	log.Fatal(http.ListenAndServe(`:`+port, server.Router))

}
