package server

import (
	"context"
	"net/http"
	"time"

	"github.com/FlyKarlik/news_skilfactory/config"
)

type Server struct {
	cfg  *config.Config
	http http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Run(handler http.Handler) error {
	s.http = http.Server{
		Addr:              s.cfg.ServerHost,
		Handler:           handler,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		ReadTimeout:       time.Minute,
	}
	return s.http.ListenAndServe()
}

func (s *Server) Shuttdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
