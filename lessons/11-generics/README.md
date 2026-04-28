# บทที่ 11: Generics

> Go เพิ่ม Generics ในเวอร์ชัน 1.18 (2022) — ช้ากว่า Java 20 ปี! แต่ design ง่ายกว่า

## Java Generics vs Go Generics

```java
// Java — type erasure
public <T extends Comparable<T>> T min(T a, T b) {
    return a.compareTo(b) <= 0 ? a : b;
}

List<String> names = new ArrayList<>();    // generic collection
```

```go
// Go — reified (type info ยังอยู่ตอน runtime)
func Min[T cmp.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

names := []string{"alice", "bob"}    // slice ไม่ต้องใส่ generic type
```

> **ความแตกต่างสำคัญ:** Go ไม่มี type erasure — generic type info ยังอยู่ตอน runtime (แต่ Go จะ monomorphize ตอน compile)

## Type Constraints

```java
// Java — ใช้ extends / super
public <T extends Number> double sum(List<T> numbers)
```

```go
// Go — ใช้ interface เป็น constraint
func Sum[T Number](numbers []T) T { ... }

type Number interface {
    ~int | ~int64 | ~float64
}
```

| Constraint | หมายถึง | Java เทียบ |
|---|---|---|
| `any` | ไม่จำกัด type | `Object` |
| `comparable` | ใช้ `==` `!=` ได้ | — |
| `cmp.Ordered` | ใช้ `<` `>` `<=` `>=` ได้ | `Comparable<T>` |
| custom interface | union types `~int \| ~float64` | `extends Number` (คร่าวๆ) |

> `~` หมายถึง "รวม type alias ด้วย" เช่น `type MyInt int` จะถูกรวมใน `~int`

## Generic Functions

```go
func Map[T any, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
// [2, 4, 6]
```

## Generic Struct

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    top := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return top, true
}
```

> **Go 1.26:** generics ใช้ได้ดีแล้ว แต่ compile ช้ากว่า code ธรรมดาเล็กน้อย

## Generic Type Aliases — 🆕 Go 1.24

```go
// Go 1.24+ — type alias มี type parameter ได้แล้ว
type Set[T comparable] = map[T]struct{}

s := make(Set[string])
s["hello"] = struct{}{}
```

> **Java เทียบ:** เหมือน `type MyList<T> = ArrayList<T>` แต่ Go ใช้ `=` (alias)

## Self-Referential Generic Constraints — 🆕 Go 1.26

```go
// Go 1.26+ — type constraint อ้างถึงตัวเองได้แล้ว
type Adder[A Adder[A]] interface {
    Add(A) A
}

func Sum[A Adder[A]](items ...A) A {
    result := items[0]
    for _, item := range items[1:] {
        result = result.Add(item)
    }
    return result
}
```

> ก่อน 1.26 จะเขียนแบบนี้ไม่ได้ — constraint ห้ามอ้างถึงตัวมันเอง

## ไฟล์ในบทนี้

- `main.go` — generic functions (Map, Filter, Reduce), generic struct (Stack), constraints
