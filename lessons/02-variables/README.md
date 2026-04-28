# บทที่ 2: Variables, Types, Constants

> Go เป็น statically typed เหมือน Java แต่เขียนสั้นกว่าเพราะมี type inference

## ประกาศตัวแปร

```java
// Java
String name = "Alice";
int age = 25;
final double PI = 3.14;
```

```go
// Go — 3 วิธี
var name string = "Alice"   // แบบเต็ม (เหมือน Java)
var name = "Alice"          // type inference
name := "Alice"             // short declaration ← ใช้บ่อยที่สุด (ใช้ได้เฉพาะใน function เท่านั้น)
```

> **สำคัญ:** `:=` ใช้ได้แค่ใน function เท่านั้น นอก function ต้องใช้ `var`

## Zero Values

Java ถ้าประกาศตัวแปรโดยไม่กำหนดค่า → compile error
Go ถ้าไม่กำหนดค่า → ได้ "zero value" ของ type นั้น

| Type | Zero Value | Java เทียบ |
|---|---|---|
| `int` | `0` | `0` |
| `float64` | `0.0` | `0.0` |
| `string` | `""` (empty string) | `null` (ถ้าเป็น Object) |
| `bool` | `false` | `false` |
| `pointer`, `slice`, `map`, `interface` | `nil` | `null` |

## Types — เทียบ Java

| Go | Java | หมายเหตุ |
|---|---|---|
| `int`, `int64` | `int`, `long` | Go มี `int` และ `int64` แยกกัน |
| `float64` | `double` | Go ใช้ `float64` เป็น default |
| `string` | `String` | Go: string เป็น **immutable byte slice** (UTF-8) |
| `bool` | `boolean` | เหมือนกัน |
| `byte` | `byte` | alias ของ `uint8` |
| `rune` | `char` (คร่าวๆ) | alias ของ `int32` — แทน Unicode code point |

## Constants

```java
// Java
public static final String APP_NAME = "MyApp";
public static final int MAX_RETRY = 3;
```

```go
// Go
const AppName = "MyApp"
const MaxRetry = 3

// หรือกลุ่ม
const (
    StatusOK    = 200
    StatusError = 500
)

// iota — auto-increment (เหมือน enum ของ Java)
const (
    Sunday    = iota  // 0
    Monday            // 1
    Tuesday           // 2
    Wednesday         // 3
)
```

> **iota** ≈ Java `enum` แต่เป็น compile-time constant แทน ไม่ใช่ object

## ไฟล์ในบทนี้

- `main.go` — ตัวอย่าง variables, types, constants, zero values
