package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {
	fmt.Println("=== 1. Default logger (TextHandler) ===")
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))
	slog.Debug("debug message", "key", "value")
	slog.Info("info message", "user", "alice", "id", 42)
	slog.Warn("warning", "temp", 99.5)
	slog.Error("error", "err", "connection refused")

	fmt.Println("\n=== 2. JSON Handler ===")
	jsonLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	jsonLogger.Info("request completed",
		"method", "GET",
		"path", "/api/users",
		"status", 200,
		"duration_ms", 42,
	)

	fmt.Println("\n=== 3. Typed attributes (slog.Attr) ===")
	slog.Info("user action",
		slog.String("action", "login"),
		slog.Int("user_id", 123),
		slog.Duration("elapsed", 150*time.Millisecond),
		slog.Bool("success", true),
	)

	fmt.Println("\n=== 4. Groups ===")
	slog.Info("http request",
		slog.Group("request",
			slog.String("method", "POST"),
			slog.String("path", "/api/orders"),
		),
		slog.Group("response",
			slog.Int("status", 201),
			slog.Duration("duration", 32*time.Millisecond),
		),
	)

	fmt.Println("\n=== 5. With — pre-populate fields ===")
	requestLogger := slog.With("request_id", "req-abc123", "service", "api")
	requestLogger.Info("processing")
	requestLogger.Info("done", "status", 200)

	fmt.Println("\n=== 6. Level control ===")
	debugLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	debugLogger.Debug("this IS visible")

	warnLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))
	warnLogger.Info("this is NOT visible")
	warnLogger.Warn("this IS visible")

	fmt.Println("\n=== 7. Custom level ===")
	customLevel := slog.Level(-8)
	slog.Log(nil, customLevel, "custom verbose level", "detail", "lots of info")

	fmt.Println("\n=== 8. MultiHandler (Go 1.26+) ===")
	multi := slog.New(slog.NewMultiHandler(
		slog.NewTextHandler(os.Stdout, nil),
	))
	multi.Info("logged via multi handler")

	fmt.Println("\n=== 9. DiscardHandler (Go 1.24+) ===")
	discard := slog.New(slog.DiscardHandler)
	discard.Info("this goes nowhere")
	fmt.Println("(discard: no output above)")

	fmt.Println("\n=== 10. Structured error logging ===")
	err := fmt.Errorf("db connection failed: %w", fmt.Errorf("timeout"))
	slog.Error("operation failed",
		"operation", "query_users",
		"error", err,
	)
}
