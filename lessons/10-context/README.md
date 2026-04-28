# บทที่ 10: Context & Concurrency Patterns

> `context.Context` คือสิ่งที่ Java ไม่มีเทียบตรง — ใช้ส่ง "สัญญาณยกเลิก", deadline, และ request-scoped data

## Context คืออะไร?

```java
// Java Spring — ใช้ RequestScope bean หรือ pass parameter manually
@RequestMapping("/api/users")
public ResponseEntity<?> getUsers(HttpServletRequest request) {
    // ไม่มี mechanism มาตรฐานสำหรับ cancellation
}
```

```go
// Go — context เป็น standard library
func handleRequest(ctx context.Context) error {
    select {
    case <-time.After(5 * time.Second):
        return nil
    case <-ctx.Done():
        return ctx.Err()   // "context canceled" หรือ "context deadline exceeded"
    }
}
```

> **context.Context** ถูกส่งผ่านทุก function ที่ทำ I/O ใน Go — เป็น convention ที่ทุกคนทำตาม

## ประเภทของ Context

| Function | หมายถึง | Java เทียบ |
|---|---|---|
| `context.Background()` | root context (ว่างเปล่า) | — |
| `context.TODO()` | "ยังไม่รู้จะใช้อะไร" | — |
| `context.WithTimeout(parent, duration)` | ยกเลิกอัตโนมัติหลัง N time | `Future.get(timeout)` |
| `context.WithCancel(parent)` | ยกเลิกได้ด้วยมือ | `Thread.interrupt()` (คร่าวๆ) |
| `context.WithValue(parent, key, value)` | ส่งข้อมูลผ่าน context | `HttpServletRequest.setAttribute()` |

## Pattern: Cancellation Propagation

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func(ctx context.Context) {
    select {
    case <-ctx.Done():
        fmt.Println("child: cancelled!")
    }
}(ctx)

cancel()   // เรียก cancel() → goroutine ที่รอ ctx.Done() จะได้สัญญาณทันที
```

## Pattern: Timeout

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

result, err := doSlowOperation(ctx)
if err == context.DeadlineExceeded {
    fmt.Println("too slow!")
}
```

## Pattern: Mutex (เมื่อต้อง share state)

```java
// Java
synchronized (lock) {
    counter++;
}
```

```go
// Go
var mu sync.Mutex
mu.Lock()
counter++
mu.Unlock()

// หรือใช้ sync.RWMutex สำหรับอ่านเยอะ เขียนน้อย
var rw sync.RWMutex
rw.RLock()    // read lock (หลายคนอ่านพร้อมกันได้)
rw.Lock()     // write lock (มีคนเดียว)
```

## ไฟล์ในบทนี้

- `main.go` — context cancel, timeout, value, mutex, goroutine + context pattern
