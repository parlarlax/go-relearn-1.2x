package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/db"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/order"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/user"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}

func main() {
	userStore := db.NewStore[*user.User]()
	userHandler := user.NewHandler(userStore)
	orderHandler := order.NewHandler()

	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)
	orderHandler.RegisterRoutes(mux)

	slog.Info("domain-driven server on :8082")
	http.ListenAndServe(":8082", loggingMiddleware(mux))
}
