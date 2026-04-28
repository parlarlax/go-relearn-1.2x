package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	slog.Info("request received",
		"method", "GET",
		"path", "/api/users",
		"status", 200,
		"duration_ms", 42,
	)

	slog.Error("database connection failed",
		"host", "localhost",
		"port", 5432,
		"err", "connection refused",
	)

	slog.Info("user action",
		slog.Group("user",
			"id", 123,
			"name", "alice",
		),
		slog.Group("action",
			"type", "login",
			"ip", "192.168.1.1",
		),
	)
}
