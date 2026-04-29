package interfaces

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	http *http.Server
	addr string
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		addr: addr,
		http: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Addr() string { return s.addr }

func (s *Server) Start() {
	slog.Info("server starting", "addr", s.addr)
	if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server error", "error", err)
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("server shutting down")
	return s.http.Shutdown(ctx)
}
