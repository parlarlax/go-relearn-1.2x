package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		http: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	slog.Info("server starting", "addr", s.http.Addr)
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("server shutting down...")
	return s.http.Shutdown(ctx)
}
