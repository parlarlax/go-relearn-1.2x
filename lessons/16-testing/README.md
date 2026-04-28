# บทที่ 16: Testing

> Go มี testing framework ในภาษาเลย — ไม่ต้องติดตั้ง JUnit/TestNG

## Java vs Go Testing

```java
// Java — JUnit 5
@Test
void testAddition() {
    assertEquals(4, Calculator.add(2, 2));
}

@Test
void testDivision() {
    assertThrows(ArithmeticException.class, () -> {
        Calculator.divide(1, 0);
    });
}
```

```go
// Go — testing package (built-in)
func TestAddition(t *testing.T) {
    got := Add(2, 2)
    want := 4
    if got != want {
        t.Errorf("Add(2,2) = %d, want %d", got, want)
    }
}

func TestDivision(t *testing.T) {
    _, err := Divide(1, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
```

## Convention

| กฎ | Java (JUnit) | Go |
|---|---|---|
| ชื่อไฟล์ | `*Test.java` | `*_test.go` |
| ชื่อ function | `@Test void testName()` | `func TestName(t *testing.T)` |
| วางที่ไหน | `src/test/java/` | **ไฟล์เดียวกัน** กับ code ปกติ |
| รัน | `mvn test` | `go test` |
| Assert | `assertEquals()`, `assertTrue()` | `if` + `t.Error()` / `t.Fatal()` |

> **สำคัญ:** Go test file วางไว้ใน package เดียวกับ code ที่ test (เช่น `user.go` + `user_test.go`)

## Table-Driven Tests — Go Pattern

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := Add(tc.a, tc.b)
            if got != tc.expected {
                t.Errorf("Add(%d,%d) = %d, want %d", tc.a, tc.b, got, tc.expected)
            }
        })
    }
}
```

> **Java เทียบ:** `@ParameterizedTest` + `@CsvSource` ใน JUnit 5 — แต่ Go ทำได้ง่ายกว่า

## Subtests — t.Run()

```go
t.Run("positive numbers", func(t *testing.T) { ... })
t.Run("negative numbers", func(t *testing.T) { ... })
```

```bash
go test -v -run TestAdd/positive    # รัน subtest เฉพาะเจาะจง
```

## Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

```bash
go test -bench=. -benchmem
```

## คำสั่งที่ใช้บ่อย

```bash
go test ./...              # รัน test ทุก package
go test -v ./...           # verbose mode
go test -run TestName      # รันเฉพาะ test ที่ชื่อตรง
go test -cover             # แสดง code coverage
go test -coverprofile=cover.out && go tool cover -html=cover.out  # coverage report
```

## `B.Loop` — เขียน Benchmark แบบใหม่ — 🆕 Go 1.24

```go
// แบบเก่า — b.N (setup รันซ้ำทุกรอบ)
func BenchmarkAddOld(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(i, i+1)
    }
}

// แบบใหม่ — b.Loop() (setup รันแค่ครั้งเดียว, ผลลัพธ์ไม่ถูก optimize ทิ้ง)
func BenchmarkAddNew(b *testing.B) {
    i := 0
    for b.Loop() {
        Add(i, i+1)
        i++
    }
}
```

> **Java เทียบ:** `b.Loop()` ≈ `@Benchmark` ใน JMH ที่ compiler ไม่ dead-code eliminate

## `T.Context()` — test context auto-cancel — 🆕 Go 1.24

```go
func TestWithContext(t *testing.T) {
    ctx := t.Context()  // auto-cancel เมื่อ test จบ
    doSomething(ctx)
}
```

## `testing/synctest` — ทดสอบ concurrent code — 🆕 Go 1.25 (GA)

```go
// synctest ช่วย test goroutine โดยจำลองเวลา (fake clock)
func TestWithSynctest(t *testing.T) {
    synctest.Test(func() {
        // goroutine ทั้งหมดในนี้ใช้ fake clock
        // time.Sleep จะข้ามทันทีเมื่อ goroutine ทุกตัว block
        go func() {
            time.Sleep(10 * time.Second)  // ไม่รอจริง!
            fmt.Println("done")
        }()
        synctest.Wait()  // รอให้ goroutine ทุกตัว block
    })
}
```

> **Java เทียบ:** `synctest` ≈ `VirtualTimeMock` ใน Spring Test + `Awaitility` แต่ built-in

## ไฟล์ในบทนี้

- `calc.go` — code ที่จะ test
- `calc_test.go` — unit tests + table-driven + benchmark
