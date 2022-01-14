package server

import (
	"fmt"
	"log"
	"net/http"
	"newExp/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	serverAdr := fmt.Sprintf("%s:%s", cfg.Http.Host, cfg.Http.Port)
	log.Println(serverAdr)
	return &Server{
		httpServer: &http.Server{
			Addr:         serverAdr,
			Handler:      handler,
			ReadTimeout:  cfg.Http.ReadTimeout,
			WriteTimeout: cfg.Http.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.httpServer.Close()
}
