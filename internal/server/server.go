package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maveonair/nut-exporter/internal/nut"
)

type Server interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type server struct {
	router    *mux.Router
	nutClient nut.Client
}

func NewServer(client nut.Client) Server {
	server := &server{
		router:    mux.NewRouter(),
		nutClient: client,
	}

	server.routes()

	return server
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
