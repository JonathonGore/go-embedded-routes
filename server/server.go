package server

import (
	"net/http"

	"github.com/JonathonGore/go-embedded-routes/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func New(api handlers.API) (*Server, error) {
	s := &Server{Router: mux.NewRouter()}

	s.Router.HandleFunc("/user", api.GetUser).Methods(http.MethodGet)
	s.Router.HandleFunc("/post", api.GetPost).Methods(http.MethodGet)

	return s, nil
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(rw, req)
}
