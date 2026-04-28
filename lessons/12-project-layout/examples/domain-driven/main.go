package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/config"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/server"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/store"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/order"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/user"
)

func main() {
	cfg := config.Load()
	slog.Info("loaded config", "port", cfg.Port, "db", cfg.DatabaseURL)

	var userStore store.UserStore = store.NewMemoryUserStore()
	var orderStore store.OrderStore = store.NewMemoryOrderStore()

	userSvc := user.NewService(userStore)
	userHandler := user.NewHandler(userSvc)

	orderSvc := order.NewService(orderStore, userStore)
	orderHandler := order.NewHandler(orderSvc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})
	userHandler.RegisterRoutes(mux)
	orderHandler.RegisterRoutes(mux)

	srv := server.New(cfg.Addr(), chain(mux, logging, recovery))

	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	slog.Info("received signal", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "error", err)
	}
	slog.Info("server stopped")
}

func chain(next http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}
	return next
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic recovered", "error", err, "path", r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"status":500,"message":"internal server error"}`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
