# บทที่ 13: Structured Logging with slog

> 🆕 **Go 1.21+** — `log/slog` เป็น standard library แล้ว ไม่ต้องใช้ logrus/zap ภายนอก

## Java Logging vs Go slog

```java
// Java — SLF4J / Logback
Logger logger = LoggerFactory.getLogger(MyService.class);
logger.info("User logged in", 
    StructuredArguments.kv("userId", 123),
    StructuredArguments.kv("ip", "192.168.1.1")
);
```

```go
// Go — slog (built-in)
import "log/slog"

slog.Info("User logged in",
    "userId", 123,
    "ip", "192.168.1.1",
)
```

## Text vs JSON Handler

```go
// Text (default) — อ่านง่ายตอน dev
slog.New(slog.NewTextHandler(os.Stdout, nil))
// time=2024-01-01T12:00:00 level=INFO msg=User logged in userId=123

// JSON — ใช้ตอน production (ส่งเข้า ELK/Datadog ได้)
slog.New(slog.NewJSONHandler(os.Stdout, nil))
// {"time":"2024-01-01T12:00:00","level":"INFO","msg":"User logged in","userId":123}
```

## Log Levels

```go
slog.Debug("debug info")     // ตั้งค่า Level >= Debug ถึงจะเห็น
slog.Info("normal log")
slog.Warn("warning")
slog.Error("error happened")
```

## slog.Group — จัดกลุ่ม key-value

```go
slog.Info("request",
    slog.Group("request",
        slog.String("method", "GET"),
        slog.String("path", "/api/users"),
    ),
    slog.Int("status", 200),
)
```

## slog.Attr — typed attributes

```go
slog.Info("user action",
    slog.String("action", "login"),
    slog.Int("userId", 42),
    slog.Duration("duration", time.Since(start)),
)
```

## Custom Handler

```go
type MyHandler struct {
    handler slog.Handler
}

func (h *MyHandler) Enabled(ctx context.Context, level slog.Level) bool {
    return h.handler.Enabled(ctx, level)
}
// ... implement Handle, WithAttrs, WithGroup
```

> **Java เทียบ:** custom handler ≈ custom Appender ใน Logback

## MultiHandler — ส่ง log หลายที่พร้อมกัน — 🆕 Go 1.26

```java
// Java — ต้อง config appender หลายตัวใน logback.xml
```

```go
// Go 1.26+ — MultiHandler (built-in)
handler := slog.NewMultiHandler(
    slog.NewJSONHandler(logFile, nil),     // เขียนลงไฟล์
    slog.NewTextHandler(os.Stdout, nil),   // พิมพ์หน้าจอ
)
slog.SetDefault(slog.New(handler))

slog.Info("this goes to both file and stdout")
```

## ไฟล์ในบทนี้

- `main.go` — text/json handler, levels, groups, typed attrs
