package main

import (
	"context"
	"log/slog"
	"os"
)

type RequestIDHandler struct {
	handler slog.Handler
}

func (h *RequestIDHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *RequestIDHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

func (h *RequestIDHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &RequestIDHandler{handler: h.handler.WithAttrs(attrs)}
}

func (h *RequestIDHandler) WithGroup(name string) slog.Handler {
	return &RequestIDHandler{handler: h.handler.WithGroup(name)}
}

func main() {
	base := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(&RequestIDHandler{handler: base})

	ctx := context.Background()
	logger.InfoContext(ctx, "with custom handler", "service", "slog-demo")
}
