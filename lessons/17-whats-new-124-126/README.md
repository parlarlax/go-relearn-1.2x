# บทที่ 17: What's New in Go 1.24 - 1.26

> ฟีเจอร์ใหม่ๆ ที่สำคัญสำหรับคนเขียน Go ในชีวิตจริง — ไม่ใช่ทุกอย่างที่ release notes บอก แต่เฉพาะที่คุณจะใช้จริง

## Runtime & Performance

### Swiss Tables Map — 🆕 Go 1.24

Map ใน Go ถูกเขียนใหม่ด้วย [Swiss Tables](https://abseil.io/about/design/swisstables) — เร็วขึ้น และใช้หน่วยความจำน้อยลง

```go
// โค้ดเดียวกัน แต่ runtime เปลี่ยน — ไม่ต้องแก้ไขอะไร
m := map[string]int{"a": 1, "b": 2}
m["c"] = 3
```

> **Java เทียบ:** เหมือน HashMap ถูก optimize ภายใน — API เดียวกัน แต่เร็วขึ้น

### Green Tea Garbage Collector — 🆕 Go 1.26 (default)

GC ใหม่ที่ลด overhead 10-40% สำหรับแอปที่ใช้ heap เยอะ

> **Java เทียบ:** เหมือนเปลี่ยนจาก Serial GC เป็น G1GC — คุณไม่ต้องทำอะไร มันดีขึ้นเอง

### Container-Aware GOMAXPROCS — 🆕 Go 1.25

```go
// ก่อน 1.25 — Go ใช้ CPU ทั้งหมดของเครื่อง แม้จะรันใน container ที่จำกัดแค่ 2 core
// 1.25+ — Go อ่าน cgroup limit อัตโนมัติ

// ใน Kubernetes: resources.limits.cpu = "2"
// → Go 1.25+ จะตั้ง GOMAXPROCS = 2 อัตโนมัติ
```

> **Java เทียบ:** เหมือน `UseContainerSupport` ใน JDK 10+ ที่อ่าน cgroup limit อัตโนมัติ

## Weak Pointers — 🆕 Go 1.24

```java
// Java — WeakReference
WeakReference<CacheEntry> ref = new WeakReference<>(entry);
CacheEntry e = ref.get();  // null ถูก GC ไปแล้ว
```

```go
// Go 1.24+ — weak package
import "weak"

type Cache[K comparable, V any] struct {
    entries map[K]weak.Pointer[V]
}

func (c *Cache[K, V]) Set(key K, value *V) {
    c.entries[key] = weak.Make(value)
}

func (c *Cache[K, V]) Get(key K) *V {
    if wp, ok := c.entries[key]; ok {
        return wp.Value()  // nil ถ้าถูก GC ไปแล้ว
    }
    return nil
}
```

> **ใช้ทำอะไร:** cache ที่ไม่กีดขวาง GC, canonicalization map, weak map

## `os.Root` — Directory Sandbox — 🆕 Go 1.24

```java
// Java — Path.resolve() ไม่มี sandbox ต้องเช็คเอง
Path basePath = Path.of("/data");
Path resolved = basePath.resolve(userInput);  // path traversal ได้ถ้า userInput = "../../etc/passwd"
```

```go
// Go 1.24+ — os.Root ป้องกัน path traversal อัตโนมัติ
root, err := os.OpenRoot("/data")
if err != nil {
    log.Fatal(err)
}

f, err := root.Open("safe.txt")         // ok
f, err := root.Open("../../etc/passwd") // ERROR! ออกนอก root ไม่ได้
f, err := root.Create("newfile.txt")    // ok
```

> **ใช้ทำอะไร:** serve static files, handle user upload, ทำอะไรกับ filesystem อย่างปลอดภัย

## `runtime.AddCleanup` — 🆕 Go 1.24

```java
// Java — Cleaner (Java 9+)
Cleaner cleaner = Cleaner.create();
cleaner.register(obj, () -> System.out.println("cleaned up!"));
```

```go
// Go 1.24+ — ดีกว่า runtime.SetFinalizer
obj := &MyResource{...}
runtime.AddCleanup(obj, func() {
    // รันเมื่อ obj ไม่ถูกใช้แล้ว
    fmt.Println("cleaned up!")
})

// ข้อดีกว่า SetFinalizer:
// - แนบได้หลาย cleanup ต่อ object เดียว
// - ไม่ทำให้เกิด memory leak จาก cycle
// - แนบได้กับ interior pointer
```

## Goroutine Leak Detection — 🆕 Go 1.26 (experimental)

```go
// เปิดด้วย GOEXPERIMENT=goroutineleakprofile
// รัน: go test -run=TestLeak ./...
// หรือดูจาก /debug/pprof/goroutineleak

func processWorkItems(ws []workItem) ([]workResult, error) {
    ch := make(chan result)
    for _, w := range ws {
        go func() {
            res, err := processWorkItem(w)
            ch <- result{res, err}  // ถ้า early return → goroutine นี้จะ leak!
        }()
    }
    for range len(ws) {
        r := <-ch
        if r.err != nil {
            return nil, r.err  // early return = goroutine ที่เหลือ leak!
        }
        results = append(results, r.res)
    }
    return results, nil
}
```

> **Java เทียบ:** เหมือน VisualVM ที่ detect thread leak แต่ทำได้อัตโนมัติผ่าน GC reachability analysis

## `go tool` Directive — 🆕 Go 1.24

```
// go.mod
tool (
    github.com/golangci/golangci-lint/cmd/golangci-lint
    golang.org/x/tools/cmd/stringer
)
```

```bash
go get tool                # อัปเดตทุก tool
go tool golangci-lint run  # รัน tool จาก go.mod
```

> **Java เทียบ:** เหมือน `pom.xml` `<build><plugins>` แต่ไม่ต้องสร้างไฟล์ `tools.go` hack แล้ว

## สรุป Version Timeline

| Version | Release | Highlights |
|---|---|---|
| 1.24 | Feb 2025 | Generic type aliases, Swiss Tables map, `weak`, `os.Root`, `B.Loop`, `T.Context`, `synctest` (exp) |
| 1.25 | Aug 2025 | Container-aware GOMAXPROCS, `synctest` GA, `WaitGroup.Go`, `json/v2` (exp), Flight Recorder |
| 1.26 | Feb 2026 | `new(expr)`, self-referential generics, Green Tea GC, goroutine leak profile, `errors.AsType`, `slog.MultiHandler` |

## ไฟล์ในบทนี้

- `main.go` — weak pointer, os.Root, AddCleanup, WaitGroup.Go examples
