# บทที่ 6: Error Handling

> Go **ไม่มี try-catch** — จัดการ error ด้วย return value แทน ซึ่งเป็นจุดที่คนจาก Java ชอกใจที่สุด

## Java vs Go — Philosophy ต่างกันมาก

```java
// Java — Exception
public String readFile(String path) throws IOException {
    BufferedReader reader = new BufferedReader(new FileReader(path));
    return reader.readLine();  // ถ้าพัง → exception กระโดดขึ้นไป
}

try {
    String content = readFile("test.txt");
} catch (IOException e) {
    System.out.println("Error: " + e.getMessage());
}
```

```go
// Go — Error as value
func readFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err   // return error ไปเลย ไม่มี throw
    }
    return string(data), nil
}

content, err := readFile("test.txt")
if err != nil {
    fmt.Println("Error:", err)
}
```

> **หลักการ Go:** Error คือ **value** ไม่ใช่ control flow mechanism — ทำให้ flow ของโปรแกรมอ่านง่าย

## Pattern: if err != nil (คุณจะเห็นทุกที่)

```go
result, err := doSomething()
if err != nil {
    return err          // ส่งต่อ
    // หรือ
    log.Fatal(err)      // หยุดโปรแกรม
    // หรือ
    return "", fmt.Errorf("context info: %w", err)  // wrap ด้วย context
}
// use result...
```

> **คุณจะเขียน `if err != nil` เยอะมาก** — นี่คือธรรมชาติของ Go ไม่ใช่ bug เป็น feature

## Custom Error

```java
// Java
public class AppException extends RuntimeException {
    private final int code;
    public AppException(int code, String msg) {
        super(msg);
        this.code = code;
    }
}
```

```go
// Go
type AppError struct {
    Code    int
    Message string
}

func (e *AppError) Error() string {
    return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func doWork() error {
    return &AppError{Code: 404, Message: "not found"}
}
```

## errors.Is() and errors.As() — เทียบ instanceof

```go
var target *AppError
if errors.As(err, &target) {
    fmt.Println("code:", target.Code)   // เข้าถึง field ของ custom error
}

if errors.Is(err, os.ErrNotExist) {
    fmt.Println("file not found")       // เช็คว่าเป็น error ตัวไหน
}
```

## Wrap/Unwrap — error chain

```go
original := errors.New("connection refused")
wrapped := fmt.Errorf("db query failed: %w", original)

fmt.Println(wrapped)              // "db query failed: connection refused"
fmt.Println(errors.Unwrap(wrapped)) // "connection refused"
fmt.Println(errors.Is(wrapped, original)) // true
```

> `%w` = wrap verb — ทำให้ error ถูก chain ไว้ สามารถ `Unwrap` กลับไปดูต้นเหตุได้

## panic/recover — "ไม่ควรใช้เป็นปกติ"

```go
// panic ≈ throw unchecked exception ใน Java
func mustFail() {
    panic("something terrible happened")
}

// recover ≈ catch (แต่ใช้กับ defer เท่านั้น)
func safeCall() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from:", r)
        }
    }()
    mustFail()
}
```

> **กฎ:** ใช้ `error` สำหรับสิ่งที่คาดไว้ได้ (file not found, network timeout) ใช้ `panic` สำหรับสิ่งที่ "เป็นไปไม่ได้" (logic bug)

## `errors.AsType` — 🆕 Go 1.26

```java
// Java — instanceof + cast
if (err instanceof AppException e) {
    System.out.println(e.getCode());
}
```

```go
// Go ก่อน 1.26 — errors.As (verbose)
var appErr *AppError
if errors.As(err, &appErr) {
    fmt.Println(appErr.Code)
}

// Go 1.26+ — errors.AsType (generic, type-safe, faster)
if appErr, ok := errors.AsType[*AppError](err); ok {
    fmt.Println(appErr.Code)
}
```

> **ข้อดี:** type-safe, เขียนสั้นกว่า, เร็วกว่า `errors.As` — ใช้ generic แทน reflection

## ไฟล์ในบทนี้

- `main.go` — error patterns, custom error, wrap/unwrap, panic/recover
