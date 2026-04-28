# บทที่ 3: Functions

> Go functions ต่างจาก Java ตรงที่รีเทิร์นได้หลายค่า และมี `defer` ซึ่ง Java ไม่มีเทียบตรง

## Function พื้นฐาน

```java
// Java
public static int add(int a, int b) {
    return a + b;
}
```

```go
// Go
func add(a int, b int) int {
    return a + b
}

// parameter type เดียวกัน เขียนย่อได้
func add(a, b int) int {
    return a + b
}
```

## Multiple Return Values — จุดเด่นของ Go

Java ถ้าอยากรีเทิร์น 2 ค่า ต้องสร้าง class/record หรือใส่ List:
```java
// Java — ต้องสร้าง wrapper class
record Result(int value, String error) {}
```

Go รีเทิร์นได้หลายค่าตรงๆ:
```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 3)
```

> นี่คือ pattern หลักของ Go แทน try-catch ของ Java — รีเทิร์น error เป็นค่าตัวเดียวกัน

## Named Return Values

```go
func swap(a, b int) (first, second int) {
    first = b
    second = a
    return  // "naked return" — return ค่าที่ชื่อไว้
}
```

> ใช้ได้แต่ไม่แนะนำให้ใช้ใน function ยาวๆ เพราะอ่านยาก

## defer — "ทำทีหลังสุด"

```java
// Java — try-with-resources
try (Connection conn = dataSource.getConnection()) {
    // use conn
} // conn.close() ถูกเรียกอัตโนมัติ
```

```go
// Go — defer
func readFile(path string) {
    file, err := os.Open(path)
    if err != nil {
        return
    }
    defer file.Close()  // ถูกเรียกตอน function จบ (ไม่ว่าจะจบยังไง)

    // use file...
}
```

> `defer` ทำงานแบบ **LIFO** (stack) — defer ตัวไหนประกาศก่อน ทำทีหลัง

## init() — เหมือน @PostConstruct

```java
// Java Spring
@PostConstruct
public void init() {
    // setup before bean is used
}
```

```go
// Go
func init() {
    // ทำงานอัตโนมัติก่อน main()
    // แต่ละ package ที่ import ก็จะมี init() ของตัวเอง
}

func main() {
    // init() ทำงานเสร็จหมดแล้ว
}
```

## ไฟล์ในบทนี้

- `main.go` — multiple return, defer, named return, init
