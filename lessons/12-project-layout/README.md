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

### แบบที่ 1: Flat (ไฟล์เดียวหรือไม่กี่ไฟล์)

```
myapp/
  go.mod
  main.go          ← ทุกอย่างอยู่ในไฟล์เดียว: handler, store, middleware
```

```java
// Java เทียบ: โยนทุกอย่างใน @RestController เดียว (Spring Boot pet project)
```

| | |
|---|---|
| **เหมาะกับ** | microservice เล็ก, CLI tool, PoC, hackathon |
| **ข้อดี** | ง่าย อ่านไว ไม่ต้องกระโดดไฟล์ |
| **ข้อเสีย** | ไม่ scale — พอเกิน ~500 บรรทัด จัดการยาก |
| **Idiomatic?** | ใช่ — Go ชอบ "start small" |

> **ตัวอย่าง:** `examples/flat/main.go`

### แบบที่ 2: Domain-Driven (แบ่งตาม feature/domain)

```
myapp/
  go.mod
  main.go              ← wiring (เชื่อมทุก domain เข้าด้วยกัน)
  user/
    user.go            ← struct + constructor
    handler.go         ← HTTP handler (เรียก store ตรงๆ)
  order/
    order.go
    handler.go
  internal/
    db/
      store.go         ← generic store (ใช้ร่วมกัน)
```

```java
// Java เทียบ: แบ่งตาม feature package
//   com.example.user/    → UserController, UserService, UserRepository
//   com.example.order/   → OrderController, OrderService, OrderRepository
```

| | |
|---|---|
| **เหมาะกับ** | แอปขนาดกลาง, REST API หลาย resource, team 2-5 คน |
| **ข้อดี** | แก้ feature หนึ่ง = แก้ใน folder เดียว, ทีมแบ่งงานตาม domain ได้ |
| **ข้อเสีย** | ต้องคิด domain boundary ให้ดี, shared code ต้องแยกเป็น `internal/` |
| **Idiomatic?** | ใช่ — นี่คือ Go style ที่แนะนำมากที่สุดสำหรับแอปขนาดกลาง |

> **ตัวอย่าง:** `examples/domain-driven/`

### แบบที่ 3: Layered (แบ่งตาม layer — แบบ Spring Boot)

```
myapp/
  go.mod
  main.go              ← wiring
  model/
    user.go            ← struct เท่านั้น
  repository/
    user_repository.go ← data access (DB)
  service/
    user_service.go    ← business logic
  handler/
    user_handler.go    ← HTTP handler
  middleware/
    logging.go         ← HTTP middleware
```

```java
// Java Spring Boot — นี่คือสไตล์ที่คุณคุ้นเคย!
//   controller/  → @RestController
//   service/     → @Service
//   repository/  → @Repository / JpaRepository
//   model/       → @Entity / DTO
```

| | |
|---|---|
| **เหมาะกับ** | คนที่มาจาก Spring Boot และอยากได้โครงสร้างคุ้นเคย, enterprise |
| **ข้อดี** | เข้าใจง่ายถ้ามาจาก Java, แยก concern ชัดเจน |
| **ข้อเสีย** | **ไม่ใช่ Idiomatic Go** — แก้ feature หนึ่ง = กระโดด 4-5 folder, overkill สำหรับแอปเล็ก |
| **Idiomatic?** | ไม่ค่อย — ใช้ได้แต่ชาว Go ส่วนใหญ่ไม่แนะนำ |

> **ตัวอย่าง:** `examples/layered/`

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
# แบบ Flat
go run ./lessons/12-project-layout/examples/flat/

# แบบ Domain-Driven
go run ./lessons/12-project-layout/examples/domain-driven/

# แบบ Layered
go run ./lessons/12-project-layout/examples/layered/

# ทดสอบ (แต่ละอันรันคนละ port: 8081, 8082, 8083)
curl http://localhost:8081/users
curl http://localhost:8082/users
curl http://localhost:8083/users
```

## ไฟล์ในบทนี้

- `main.go` — ตัวอย่าง UserService แบบ simple
- `examples/flat/` — Flat structure: ทุกอย่างใน main.go
- `examples/domain-driven/` — Domain-Driven: user/, order/, internal/db/
- `examples/layered/` — Layered: model/, repository/, service/, handler/, middleware/
