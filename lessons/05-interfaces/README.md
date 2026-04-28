# บทที่ 5: Interface & Composition

> Go interface ต่างจาก Java interface ตรงที่ **ไม่ต้อง implements** — implement แบบ implicit (duck typing)

## Java vs Go Interface

```java
// Java — ต้องระบุ implements อย่างชัดเจน
public interface Speaker {
    String speak();
}

public class Dog implements Speaker {
    public String speak() { return "Woof"; }
}
```

```go
// Go — ไม่ต้องระบุ implements แค่มี method ตรงกันก็พอ
type Speaker interface {
    Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string {
    return d.Name + " says Woof!"
}

// Dog ถูกต้องกฎ Speaker อัตโนมัติ — ไม่ต้องเขียน "implements Speaker"
```

> **กฎ:** ถ้า struct มี method signature ตรงกับ interface ทุกอัน → struct นั้น implement interface นั้นอัตโนมัติ (Structural Typing / Duck Typing)

## ทำไมถึงดีกว่า?

- **Decoupling:** interface และ implementation ไม่ต้องรู้จักกัน
- **Test:** ง่ายต่อการ mock สำหรับ test
- **Flexibility:** ใครๆ ก็ implement interface ของคุณได้ แม้แต่ code ที่เขียนทีหลัง

## Interface Composition

```java
// Java — ต้อง extends interface อื่น
public interface ReadWriter extends Reader, Writer {}
```

```go
// Go — embed interface เข้าด้วยกัน
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

## Empty Interface — `interface{}` (เหมือน `Object` ของ Java)

```go
func PrintAny(v interface{}) {    // รับอะไรก็ได้
    fmt.Println(v)
}

// Go 1.18+ แนะนำใช้ any แทน interface{}
func PrintAny(v any) {            // any = alias ของ interface{}
    fmt.Println(v)
}
```

> **Java เทียบ:** `Object` — ทุกอย่างเป็น Object ใน Java, ทุกอย่างเป็น `any` ใน Go

## Type Assertion & Type Switch

```java
// Java — instanceof
if (obj instanceof String s) {
    System.out.println(s.length());
}
```

```go
// Go — type assertion
str, ok := v.(string)    // ok = true ถ้าแปลงได้

// type switch
switch v := val.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown:", v)
}
```

## Nil Interface Trap — ระวัง!

```go
var err error = nil      // nil interface — ok
var e *MyError = nil     // typed nil pointer
err = e                  // err ตอนนี้ != nil !!!

if err != nil {
    // เข้าประตูนี้! เพราะ err มี type (*MyError) อยู่ แม้ค่าเป็น nil
}
```

> **สาเหตุ:** Go interface มี 2 ส่วน — (type, value) — ถ้า type ไม่เป็น nil ถือว่า interface ไม่เป็น nil

## `new()` with Expression — 🆕 Go 1.26

```java
// Java — constructor รับค่าได้เลย
Integer age = Integer.valueOf(25);
```

```go
// Go ก่อน 1.26 — ต้องสร้างตัวแปรก่อน
ageVal := 25
age := &ageVal

// Go 1.26+ — new() รับ expression ได้เลย!
age := new(25)

// เหมาะมากกับ optional JSON fields
type Person struct {
    Name string `json:"name"`
    Age  *int   `json:"age"`   // pointer = optional
}

p := Person{
    Name: "Alice",
    Age:  new(25),             // กรอก optional field ได้ในบรรทัดเดียว!
}
```

## ไฟล์ในบทนี้

- `main.go` — interface, composition, type switch, nil trap
