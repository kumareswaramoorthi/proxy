package server

import (
	"net/http"
	"proxy/handler"
	"proxy/middleware"
)

type Server struct {
	Router *http.ServeMux
}

func (s *Server) InitRoute(h *handler.Handler) {
	s.Router.HandleFunc("/microservice/name", middleware.JSONandCORS(h.MicroserviceName))
	s.Router.HandleFunc("/user/name", middleware.JSONandCORS(h.User))
}
