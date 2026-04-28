# บทที่ 1: Hello Go

> ถ้าคุณเคยเขียน Java Spring Boot มา ให้คิดว่า Go คือเวอร์ชันที่ "ไม่มี ceremony" — ไม่มี class, ไม่มี annotation, ไม่ต้อง build นาน

## Java vs Go — ภาพรวม

|  concept | Java / Spring Boot | Go |
|---|---|---|
| Entry point | `public static void main(String[] args)` | `func main()` |
| Build tool | Maven / Gradle (`pom.xml` / `build.gradle`) | `go` command (`go.mod`) |
| Package manager | Maven Central | Go Modules (proxy.golang.org) |
| Compile | `javac` → `.class` → JAR | `go build` → binary เดียว |
| Run | `java -jar app.jar` (ต้องมี JVM) | `./app` (รันได้เลย ไม่ต้องมี runtime) |

## สร้างโปรเจกต์

```bash
mkdir myapp && cd myapp
go mod init github.com/yourname/myapp    # เหมือน pom.xml ของ Maven
```

ไฟล์ `go.mod` จะถูกสร้างให้ — มันคือ "identity" ของ module คุณ

## โค้ด Go แรก

เปรียบเทียบกับ Java:

```java
// Java — Hello.java
public class Hello {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

```go
// Go — main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

สังเกต:
- **ไม่ต้องสร้าง class** — Go ใช้ `package` + `func` ตรงๆ
- `fmt` = `System.out` ของ Go (format + print)
- **ไม่มี semicolon** — Go compiler ใส่ให้อัตโนมัติ

## รัน

```bash
go run main.go        # compile + run ทีเดียว (เหมือน java Hello.java ใน Java 11+)
go build              # compile เป็น binary (เหมือน mvn package)
./myapp               # รัน binary ที่ build แล้ว
```

## fmt package — การจัดรูปแบบข้อความ

```go
fmt.Println("Hello")                    // พิมพ์พร้อมขึ้นบรรทัดใหม่
fmt.Printf("Name: %s, Age: %d\n", ... ) // format string (เหมือน String.format ของ Java)
fmt.Sprintf("Hi %s", name)              // ส่งคืน string (ไม่พิมพ์)
```

| Verb | หมายถึง | Java เทียบ |
|---|---|---|
| `%s` | string | `%s` |
| `%d` | integer | `%d` |
| `%f` | float | `%f` |
| `%v` | ค่า default (ใช้ได้กับทุก type) | `toString()` |
| `%+v` | พร้อม field name | — |
| `%T` | แสดง type | `obj.getClass()` |

## ไฟล์ในบทนี้

- `main.go` — โค้ดตัวอย่าง Hello Go + fmt
