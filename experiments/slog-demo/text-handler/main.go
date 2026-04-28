package main

import (
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	slog.Debug("debug message", "key", "value")
	slog.Info("info message", "user", "alice", "id", 42)
	slog.Warn("warning message", "temp", 99.5)
	slog.Error("error message", "err", "something failed")
}
