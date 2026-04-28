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

## ไฟล์ในบทนี้

- `main.go` — ตัวอย่างการแบ่ง package และ import
