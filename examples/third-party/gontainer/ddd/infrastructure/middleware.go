package infrastructure

import (
	"log/slog"
	"net/http"
	"time"
)

type HTTPMiddleware interface {
	Wrap(http.Handler) http.Handler
}

type LoggingMiddleware struct{}

func NewLoggingMiddleware() *LoggingMiddleware { return &LoggingMiddleware{} }

func (m *LoggingMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}

type RecoveryMiddleware struct{}

func NewRecoveryMiddleware() *RecoveryMiddleware { return &RecoveryMiddleware{} }

func (m *RecoveryMiddleware) Wrap(next http.Handler) http.Handler {
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
