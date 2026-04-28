# บทที่ 12: Packages & Project Layout

> Java ใช้ package + Maven/Gradle + Spring Boot structure — Go ใช้ module + flat structure

## Java vs Go — Package System

|  concept | Java | Go |
|---|---|---|
| Package | `com.example.app` (hierarchical) | `github.com/yourname/app` (flat URL-based) |
| Build file | `pom.xml` / `build.gradle` | `go.mod` |
| Dependency | Maven Central | Go Modules Proxy |
| Private | `private` keyword | ตัวพิมพ์เล็ก |
| Internal | (ไม่มี built-in) | `internal/` folder |

## go.mod — เทียบ pom.xml

```
module github.com/yourname/myapp

go 1.26

require (
    github.com/some/dependency v1.2.3
)
```

```bash
go mod init github.com/yourname/myapp   # สร้าง go.mod
go mod tidy                              # ดาวน์โหลด dep + ล้าง unused
go get github.com/some/pkg@v1.2.3        # เพิ่ม dependency
```

## Idiomatic Go Layout (เทียบ Spring Boot)

```
# Spring Boot                      # Go
src/main/java/com/example/         → root (หรือชื่อ module)
  controller/                      → ไม่ต้องแยก! ใช้ flat package
  service/                         → ไม่ต้องแยก! logic อยู่ใน package เดียวกัน
  repository/                      → ไม่ต้องแยก! storage logic อยู่ใน package เดียวกัน
  model/                           → ไม่ต้องแยก! struct อยู่ใน package เดียวกัน
  config/                          → ไม่ต้องแยก! config อยู่ใน package เดียวกัน
```

```
# Go layout (แบบ idiomatic)

myapp/
  go.mod
  main.go              ← entry point
  user.go              ← struct + methods + handler (ทั้งหมดใน package เดียวกัน!)
  user_test.go
  db.go
  internal/            ← ใช้เมื่อต้องการซ่อน code จากภายนอก
    auth/
      auth.go
  cmd/
    server/
      main.go          ← entry point แยก (ถ้ามีหลาย binary)
```

## internal/ — Go เทียบ package-private (แต่แรงกว่า)

```java
// Java — package-private (คนนอก package เดียวกันเข้าไม่ได้)
class InternalHelper { ... }
```

```go
// Go — internal/ (คนนอก module เรียกใช้ไม่ได้เลย — compiler บังคับ)
// anything under internal/ ถูก import ได้เฉพาะภายใน module เดียวกันเท่านั้น
```

## กฎง่ายๆ ของ Go Layout

1. **เริ่มจาก flat** — ไฟล์ `.go` ที่ root เลย ถ้าโปรเจกต์ยังเล็ก
2. **แยกเมื่อซับซ้อน** — แยกตาม domain (`user/`, `order/`, `payment/`) ไม่ใช่ตาม layer (`controller/`, `service/`)
3. **cmd/** — เก็บ entry point ถ้ามีหลาย binary
4. **internal/** — ใช้เมื่อต้องการ encapsulation
5. **อย่าสร้าง pkg/** — ใน Go ยุค modern ไม่จำเป็น

> อ่านเพิ่มเติมใน Q&A.md หัวข้อ "golang-standards/project-layout"

## Go CLI Reference — คำสั่งที่ใช้บ่อย

> เทียบกับ Java (Maven/Gradle) เพื่อให้เทียบง่าย

### Build & Run

| Go | Java (Maven) | หมายถึง |
|---|---|---|
| `go run main.go` | `java Main.java` (Java 11+) | compile + run ทีเดียว |
| `go run ./...` | — | run ทุก package |
| `go build` | `mvn compile` | compile (สร้าง binary ถ้าเป็น main package) |
| `go build -o myapp` | `mvn package` | compile + ตั้งชื่อ output |
| `go install` | `mvn install` | compile + ติดตั้งไป `$GOBIN` |

### Dependencies & Modules

| Go | Java (Maven) | หมายถึง |
|---|---|---|
| `go mod init <module>` | สร้าง `pom.xml` | สร้าง module ใหม่ |
| `go mod tidy` | `mvn dependency:resolve` | ดาวน์โหลด deps + ลบ unused |
| `go mod download` | `mvn dependency:go-offline` | ดาวน์โหลดเท่านั้น |
| `go get pkg@v1.2.3` | เพิ่มใน pom.xml + `mvn install` | เพิ่ม/อัปเกรด dependency |
| `go get pkg@latest` | `mvn versions:use-latest` | อัปเกรดเป็นเวอร์ชันล่าสุด |

### Test

| Go | Java (JUnit) | หมายถึง |
|---|---|---|
| `go test ./...` | `mvn test` | รัน test ทุก package |
| `go test -v` | `mvn test` (verbose) | แสดงผลละเอียด |
| `go test -run TestName` | รัน test เฉพาะ method | filter test |
| `go test -cover` | JaCoCo plugin | แสดง code coverage |
| `go test -bench=.` | JMH benchmark | รัน benchmark |
| `go test -race` | — | detect race condition |

### Code Quality

| Go | Java | หมายถึง |
|---|---|---|
| `go vet ./...` | SpotBugs / Error Prone | ตรวจจับ bug ทั่วไป (ผิด fmt, unreachable code, ฯลฯ) |
| `go fmt` | `google-java-format` | จัด format โค้ดให้ตรงมาตรฐาน |
| `go doc fmt.Println` | `javadoc` | ดู documentation |
| `go fix` | — | 🆕 Go 1.26: อัปเดตโค้ดเป็น idiom ใหม่อัตโนมัติ |

### คำสั่งอื่นๆ ที่มีประโยชน์

```bash
go version                    # เช็คเวอร์ชัน Go
go env                        # ดู environment variables
go env GOPATH                 # ดูค่าเฉพาะตัว
go list -m all                # ลิสต์ dependencies ทั้งหมด
go mod graph                  # แสดง dependency tree (เหมือน mvn dependency:tree)
go mod why github.com/pkg     # อธิบายว่าทำไมถึงต้องใช้ pkg นี้
go tool cover -html=cover.out # เปิด coverage report ใน browser
```

> **Tip:** ใช้ `go help <command>` เพื่อดู help ของแต่ละคำสั่ง เช่น `go help test`

---

## เปรียบเทียบ 3 สไตล์ Project Structure

ทั้ง 3 แบบทำงานเหมือนกัน (CRUD API) แต่จัดโครงสร้างต่างกัน — ดูโค้ดตัวอย่างได้ใน `examples/`

### แบบที่ 1: Flat (ไฟล์เดียว — ทุกอย่างอยู่ด้วยกัน)

```
myapp/
  go.mod
  main.go          ← interface, store, service, handler, middleware, server — ทั้งหมดในไฟล์เดียว
```

```java
// Java เทียบ: โยนทุกอย่างใน @RestController เดียว (Spring Boot pet project)
```

#### สิ่งที่อยู่ในไฟล์เดียว (แต่แยก section ชัดเจน):

```go
// 1. Models
type User struct { ... }
type Order struct { ... }

// 2. Store interfaces — สามารถ swap เป็น postgres ได้
type UserStore interface { ... }

// 3. Store implementation (in-memory)
type MemoryUserStore struct { ... }

// 4. Business logic (validation)
type UserService struct { store UserStore }
func NewUserService(store UserStore) *UserService { ... }

// 5. HTTP handlers
func registerUserRoutes(mux *http.ServeMux, svc *UserService) { ... }

// 6. Middleware
func loggingMiddleware(next http.Handler) http.Handler { ... }
func recoveryMiddleware(next http.Handler) http.Handler { ... }

// 7. Wiring + graceful shutdown
func main() {
    var userStore UserStore = NewMemoryUserStore()      // DI
    userSvc := NewUserService(userStore)                 // DI
    mux := http.NewServeMux()
    registerUserRoutes(mux, userSvc)                     // routes
    handler := chainMiddleware(mux, loggingMiddleware, recoveryMiddleware) // middleware
    // ... graceful shutdown
}
```

> **สังเกต:** logic เหมือน domain-driven และ layered ทุกอย่าง — แค่ไม่แยกไฟล์
> interface + DI + validation + middleware + graceful shutdown มีครบเหมือนกัน

| | |
|---|---|
| **เหมาะกับ** | microservice เล็ก, CLI tool, PoC, hackathon |
| **ข้อดี** | ง่าย อ่านไว ไม่ต้องกระโดดไฟล์ |
| **ข้อเสีย** | ไม่ scale — พอเกิน ~500 บรรทัด จัดการยาก |
| **Idiomatic?** | ใช่ — Go ชอบ "start small" |

> **ตัวอย่าง:** `examples/flat/main.go`

### แบบที่ 2: Domain-Driven (แบ่งตาม feature/domain) — แนะนำ!

```
myapp/
  go.mod
  main.go                        ← wiring: สร้าง deps, เชื่อมทุกอย่างเข้าด้วยกัน
  user/
    service.go                   ← business logic (รับ interface ไม่ใช่ concrete)
    handler.go                   ← HTTP handler (รับ Service)
  order/
    service.go                   ← business logic (รับ UserStore + OrderStore)
    handler.go                   ← HTTP handler
  internal/
    config/
      config.go                  ← อ่าน config จาก env vars
    server/
      server.go                  ← HTTP server wrapper + graceful shutdown
    store/
      model.go                   ← struct definitions + interfaces
      memory.go                  ← in-memory implementation (swap เป็น SQL ได้)
```

```java
// Java เทียบ: แบ่งตาม feature package
//   com.example.user/    → UserController, UserService, UserRepository
//   com.example.order/   → OrderController, OrderService, OrderRepository
```

#### Real-World Patterns ที่ใช้ในตัวอย่างนี้:

**1. Dependency Injection — Go vs Spring**

```java
// Spring Boot — framework ทำให้ (magic!)
@Service
public class UserService {
    @Autowired                    // ← framework inject ให้อัตโนมัติ
    private UserRepository repo;
}
```

```go
// Go — constructor injection (manual, explicit)
type Service struct {
    users store.UserStore         // ← รับ interface ไม่ใช่ concrete type
}

func NewService(users store.UserStore) *Service {
    return &Service{users: users} // ← inject ด้วยมือตอนสร้าง
}
```

> **Go DI = explicit** — ไม่มี `@Autowired`, ไม่มี magic, ทุก dependency ปรากฏใน function signature หมด

**2. Interface-Based Store — เปลี่ยน implementation ได้ไม่ต้องแก้ code**

```go
// store/model.go — นิยาม interface
type UserStore interface {
    Create(name, email string) (*User, error)
    GetByID(id int) (*User, error)
    List() ([]*User, error)
    Update(id int, name, email string) (*User, error)
    Delete(id int) error
}

// store/memory.go — implement ด้วย map (สำหรับ dev/test)
// store/postgres.go — implement ด้วย database/sql (สำหรับ production)
```

```java
// Java เทียบ: interface + @Repository + @Autowired
//   แต่ Go ไม่ต้องมี annotation หรือ IoC container
```

**3. main.go as Wiring — เหมือน Spring AppConfig**

```go
func main() {
    cfg := config.Load()                                    // 1. อ่าน config

    var userStore store.UserStore = store.NewMemoryUserStore()  // 2. สร้าง deps
    var orderStore store.OrderStore = store.NewMemoryOrderStore()

    userSvc := user.NewService(userStore)                       // 3. inject deps
    userHandler := user.NewHandler(userSvc)

    orderSvc := order.NewService(orderStore, userStore)         // 4. order ต้องการ user ด้วย
    orderHandler := order.NewHandler(orderSvc)

    mux := http.NewServeMux()
    userHandler.RegisterRoutes(mux)                             // 5. register routes
    orderHandler.RegisterRoutes(mux)

    srv := server.New(cfg.Addr(), chain(mux, logging, recovery)) // 6. middleware chain
    srv.Start()                                                 // 7. start + graceful shutdown
}
```

> **Java เทียบ:** `main.go` = `@SpringBootApplication` + `@Bean` methods รวมไว้ในที่เดียว

**4. Middleware Chain**

```go
func chain(next http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        next = middlewares[i](next)
    }
    return next
}

// ใช้: chain(mux, logging, recovery)
// request → logging → recovery → actual handler
```

**5. Graceful Shutdown**

```go
go func() {
    srv.Start()            // รัน server ใน goroutine
}()

quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit                     // รอ SIGINT (Ctrl+C)

srv.Shutdown(ctx)          // รอ request ที่ค้างอยู่ให้จบก่อน
```

| | |
|---|---|
| **เหมาะกับ** | แอปขนาดกลาง, REST API หลาย resource, team 2-5 คน |
| **ข้อดี** | แก้ feature หนึ่ง = แก้ใน folder เดียว, swap store ได้, test ง่าย |
| **ข้อเสีย** | ต้องคิด domain boundary ให้ดี, wiring code ใน main ยาวขึ้นเรื่อยๆ |
| **Idiomatic?** | ใช่ — นี่คือ Go style ที่แนะนำมากที่สุด |

> **ตัวอย่าง:** `examples/domain-driven/`

### แบบที่ 3: Layered (แบ่งตาม layer — แบบ Spring Boot)

```
myapp/
  go.mod
  main.go                    ← wiring: สร้าง deps layer by layer
  config/
    config.go                 ← อ่าน config จาก env vars
  model/
    model.go                  ← struct definitions (เทียบ @Entity / DTO)
  repository/
    repository.go             ← interface + memory implementation (เทียบ @Repository)
  service/
    service.go                ← business logic + validation (เทียบ @Service)
  handler/
    handler.go                ← HTTP handlers (เทียบ @RestController)
  middleware/
    middleware.go              ← logging, recovery, chain
  server/
    server.go                 ← HTTP server wrapper + graceful shutdown
```

```java
// Java Spring Boot — นี่คือสไตล์ที่คุณคุ้นเคย!
//   controller/  → @RestController
//   service/     → @Service
//   repository/  → @Repository / JpaRepository
//   model/       → @Entity / DTO
//   config/      → @Configuration
```

#### Wiring: main.go (Layer by Layer)

```go
func main() {
    cfg := config.Load()

    // Layer 1: Repository (Data Access)
    userRepo := repository.NewMemoryUserRepository()
    orderRepo := repository.NewMemoryOrderRepository()

    // Layer 2: Service (Business Logic) — รับ interface ไม่ใช่ concrete
    userSvc := service.NewUserService(userRepo)
    orderSvc := service.NewOrderService(orderRepo, userRepo)  // cross-domain dep

    // Layer 3: Handler (HTTP) — รับ Service
    userHandler := handler.NewUserHandler(userSvc)
    orderHandler := handler.NewOrderHandler(orderSvc)

    // Layer 4: Routes + Middleware
    mux := http.NewServeMux()
    userHandler.RegisterRoutes(mux)
    orderHandler.RegisterRoutes(mux)

    srv := server.New(cfg.Addr(), middleware.Chain(mux, middleware.Logging, middleware.Recovery))
    srv.Start()  // + graceful shutdown
}
```

> **Java เทียบ:** ทุกบรรทัดข้างบน = สิ่งที่ Spring IoC Container ทำให้ด้วย `@Autowired` อัตโนมัติ
> ใน Go คุณเขียน "wiring" เอง — ยาวกว่าแต่ **อ่านรู้เรื่องทุกบรรทัด**

#### ข้อสังเกต: เปรียบเทียบกับ Domain-Driven

| | Layered (แบบนี้) | Domain-Driven (แบบที่แล้ว) |
|---|---|---|
| แก้ feature "user" | แก้ **4 folder**: model/ + repository/ + service/ + handler/ | แก้ **1 folder**: user/ |
| เพิ่ม domain "payment" | ต้องแก้ **ทุก layer** | สร้าง **folder เดียว**: payment/ |
| Dependency flow | `handler → service → repository → model` (เดียวกัน) | `handler → service → store` (เดียวกัน) |
| DI pattern | constructor injection (เหมือนกัน) | constructor injection (เหมือนกัน) |

> **สรุป:** Logic เหมือนกันทุกอย่าง ต่างแค่ว่าไฟล์อยู่ "รวมตาม domain" หรือ "แยกตาม layer"

### สรุป: ใช้แบบไหนดี?

| ขนาดโปรเจกต์ | แนะนำ | เหตุผล |
|---|---|---|
| เล็ก (< 1000 บรรทัด) | **Flat** | ไม่ต้องซับซ้อน |
| กลาง (API หลาย resource) | **Domain-Driven** | แก้ทีละ feature สะดวก |
| ใหญ่ / Enterprise / Team ใหญ่ | **Domain-Driven** + `internal/` | domain boundary ชัด + ซ่อน implementation |
| ยังไม่รู้ | **Flat → Domain** | เริ่ม flat แล้ว refactor เมื่อซับซ้อน |

> **กฎทอง:** เริ่มจาก flat เสมอ แล้วแยกเมื่อรู้สึกว่าไฟล์เดียวมันเยอะเกิน — ไม่ใช่แยกก่อนแล้วค่อยเขียน

## รันตัวอย่าง

```bash
# แบบ Flat (port 8081)
go run ./lessons/12-project-layout/examples/flat/

# แบบ Domain-Driven (port 8080) — แนะนำ!
go run ./lessons/12-project-layout/examples/domain-driven/

# แบบ Layered (port 8083)
go run ./lessons/12-project-layout/examples/layered/

# ทดสอบ Domain-Driven (real-world example)
curl http://localhost:8080/users
curl http://localhost:8080/users/1
curl -X POST http://localhost:8080/users -d '{"name":"Alice","email":"alice@test.com"}'
curl -X PUT http://localhost:8080/users/1 -d '{"name":"Updated","email":"new@test.com"}'
curl -X DELETE http://localhost:8080/users/1
curl -X POST http://localhost:8080/orders -d '{"user_id":1,"item":"Go book","qty":2}'
curl http://localhost:8080/users/1/orders
```

## ไฟล์ในบทนี้

- `main.go` — ตัวอย่าง UserService แบบ simple
- `examples/flat/` — Flat structure: ทุกอย่างใน main.go
- `examples/domain-driven/` — Domain-Driven: user/, order/, internal/db/
- `examples/layered/` — Layered: model/, repository/, service/, handler/, middleware/
