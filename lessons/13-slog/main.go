package main

import (
	"fmt"
	"log/slog"
	"os"
)

func textHandlerDemo() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	logger.Debug("debug: only visible when level >= Debug")
	logger.Info("info: user action", "userId", 42, "action", "login")
	logger.Warn("warn: high temperature", "temp", 99.5)
	logger.Error("error: connection failed", "host", "localhost")
}

func jsonHandlerDemo() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("request completed",
		"method", "GET",
		"path", "/api/users/123",
		"status", 200,
		"duration_ms", 42,
	)

	logger.Info("user action",
		slog.Group("user",
			slog.Int("id", 123),
			slog.String("name", "alice"),
		),
		slog.Group("request",
			slog.String("method", "POST"),
			slog.String("ip", "192.168.1.1"),
		),
	)
}

func defaultLogger() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	slog.Info("using default logger", "service", "slog-demo")
}

func main() {
	fmt.Println("=== text handler ===")
	textHandlerDemo()

	fmt.Println("\n=== json handler ===")
	jsonHandlerDemo()

	fmt.Println("\n=== default logger ===")
	defaultLogger()
}
