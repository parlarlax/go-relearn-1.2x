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

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/config"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/handler"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/middleware"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/repository"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/server"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/service"
)

func main() {
	cfg := config.Load()
	slog.Info("loaded config", "port", cfg.Port, "db", cfg.DatabaseURL)

	userRepo := repository.NewMemoryUserRepository()
	orderRepo := repository.NewMemoryOrderRepository()

	userSvc := service.NewUserService(userRepo)
	orderSvc := service.NewOrderService(orderRepo, userRepo)

	userHandler := handler.NewUserHandler(userSvc)
	orderHandler := handler.NewOrderHandler(orderSvc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})
	userHandler.RegisterRoutes(mux)
	orderHandler.RegisterRoutes(mux)

	srv := server.New(cfg.Addr(), middleware.Chain(mux, middleware.Logging, middleware.Recovery))

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
