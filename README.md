# Go Re-learning Lab — Zero to Hero (Java Spring Boot → Go 1.26)

Repository สำหรับเรียน Go ตั้งแต่ศูนย์ ออกแบบมาสำหรับคนที่เคยเขียน **Java / Spring Boot** มา

## โครงสร้าง

```
lessons/
  01-hello-go/       → Hello Go, fmt, go mod
  02-variables/      → Variables, Types, Constants, Zero Values
  03-functions/      → Multiple return, defer, init
  04-struct-methods/ → Struct, Methods, Embedding (composition)
  05-interfaces/     → Interface, Duck Typing, Type Switch, Nil Trap, new() expr (1.26)
  06-error-handling/ → error pattern, custom error, wrap/unwrap, panic/recover, errors.AsType (1.26)
  07-slices-maps/    → Slice, Map, Range, make, copy
  08-pointers/       → Pointer basics, value vs reference
  09-goroutines/     → Goroutine, Channel, WaitGroup, Select
  10-context/        → Context cancel/timeout/value, Mutex
  11-generics/       → Generic functions, constraints, generic struct, type aliases (1.24), self-ref (1.26)
  12-project-layout/ → Package layout, Idiomatic Go structure
  13-slog/           → Structured Logging (Go 1.21+), MultiHandler (1.26)
  14-http-server/    → HTTP Server + Router (Go 1.22+)
  15-iterators/      → Range-over-func Iterators (Go 1.23+)
  16-testing/        → Unit Test, Table-Driven, Benchmark, B.Loop (1.24), synctest (1.25)
  17-whats-new-124-126/ → weak, os.Root, AddCleanup, WaitGroup.Go, GC, goroutine leak
```

## วิธีรันแต่ละบท

```bash
go run ./lessons/01-hello-go/
go run ./lessons/02-variables/
go run ./lessons/09-goroutines/
# ... และอื่นๆ

# บท 14 (HTTP Server) ต้องทดสอบด้วย curl:
go run ./lessons/14-http-server/
curl http://localhost:8080/users

# บท 16 (Testing):
go test -v ./lessons/16-testing/
```

## Version Notes

| บท | ฟีเจอร์ | Go Version |
|---|---|---|
| 01-12 | พื้นฐานภาษา Go | 1.0+ |
| 11 | Generics | **1.18+** |
| 11 | Generic type aliases | **1.24+** |
| 13 | `log/slog` structured logging | **1.21+** |
| 14 | `net/http` method + path params | **1.22+** |
| 15 | `range-over-func` iterators | **1.23+** |
| 16 | `B.Loop`, `T.Context`, synctest | **1.24+** |
| 17 | weak, os.Root, Swiss Tables, GC | **1.24-1.26** |

## Examples

```
examples/
├── stdlib/          ← Go standard library examples (one topic per directory)
└── third-party/     ← Third-party library examples (gontainer DI, etc.)
```

### Standard Library

```bash
go run ./examples/stdlib/slog/
go run ./examples/stdlib/encoding-json/
go run ./examples/stdlib/database-sql/
# ... และอื่นๆ
```

| Directory | สอนเรื่อง | Java Equivalent |
|---|---|---|
| `stdlib/slog/` | Structured logging (Text, JSON, Groups, MultiHandler) | SLF4J + Logback |
| `stdlib/encoding-json/` | Marshal, Unmarshal, omitzero, RawMessage, custom JSON | Jackson ObjectMapper |
| `stdlib/slices-maps/` | Sort, Contains, Insert, Delete, BinarySearch, Concat | Stream API, Collections |
| `stdlib/strings-bytes/` | Split, Join, Builder, iterators (1.24+) | String utilities |
| `stdlib/context/` | WithTimeout, WithCancel, WithValue, Deadline | ThreadLocal / RequestScope |
| `stdlib/sync/` | Mutex, WaitGroup.Go, Once, Cond, Pool, Map | synchronized, CountDownLatch, ConcurrentHashMap |
| `stdlib/time/` | Format, Parse, Duration, Timer, Ticker, AfterFunc | java.time, ScheduledExecutor |
| `stdlib/crypto/` | crypto/rand, sha256, HKDF (1.24+) | java.security, javax.crypto |
| `stdlib/iter/` | Backward, Filter, Map, Take, BinaryTree iterator | Stream API |
| `stdlib/os/` | ReadFile, WriteFile, MkdirAll, env, Args, os.Root | java.nio.file |
| `stdlib/unique/` | Interning / canonicalization (1.23+) | String.intern() / ConcurrentHashMap cache |
| `stdlib/regexp/` | Compile, Find, groups, Replace, Split | java.util.regex |
| `stdlib/encoding-xml/` | Marshal, Unmarshal, XML tags, Encoder, CDATA | JAXB |
| `stdlib/testing/` | Table-driven, error testing, benchmark, B.Loop (1.24) | JUnit 5 |
| `stdlib/net-http/` | Server routes + client, path params, headers | Spring MVC + RestTemplate |
| `stdlib/database-sql/` | CRUD, transactions, prepared stmt, connection pool (pure-Go SQLite) | JdbcTemplate + HikariCP |

### Third-Party Libraries

| Directory | Library | Description | Java Equivalent |
|---|---|---|---|
| `third-party/gontainer/flat/` | [NVIDIA/gontainer](https://github.com/NVIDIA/gontainer) v2 | DI basics: factory, cleanup, Optional, Multiple, transient, annotations | Spring IoC Container |
| `third-party/gontainer/layer/` | [NVIDIA/gontainer](https://github.com/NVIDIA/gontainer) v2 | Layered architecture: handler → service → repository → model | Spring Boot layered |
| `third-party/gontainer/ddd/` | [NVIDIA/gontainer](https://github.com/NVIDIA/gontainer) v2 | Domain-Driven Design: domain/application/infrastructure/interfaces | Spring Boot DDD |

## Prerequisites

- Go 1.26+ (`go.dev/dl`)
- (Optional) VS Code + Go Extension
