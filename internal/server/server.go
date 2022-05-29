package server

import (
	"fmt"
	"hemli/internal/router"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *router.Router
	// Db, config can be added here
}

func CreateNewServer() *Server {
	return &Server{
		router: router.NewRouter(chi.NewRouter()),
	}
}

func (s *Server) Run(host, port string) {
	s.router.MountHandlers()

	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Application started on: http://%s\n", addr)
	http.ListenAndServe(addr, s.router.Router)
}
