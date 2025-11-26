package server

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handlers interface {
	Roll(w http.ResponseWriter, r *http.Request)
}
type Server struct {
	handlers handlers
	server   *http.Server
	router   *chi.Mux
}

func NewServer(handl handlers) *Server {
	return &Server{
		handlers: handl,
		router:   chi.NewRouter(),
	}
}

func (s *Server) Start() error {

	port := ":8080"

	s.router.Route("/dice", func(r chi.Router) {
		r.Post("/", s.handlers.Roll)

	})

	s.server = &http.Server{
		Addr:    port,
		Handler: s.router,
	}
	log.Printf("started")
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
