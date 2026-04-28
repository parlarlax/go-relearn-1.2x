# บทที่ 15: Iterators — range-over-func

> 🆕 **Go 1.23+** — Go รองรับ custom iterator ผ่าน `range-over-func` — ทำให้ `for range` ใช้กับอะไรก็ได้

## Iterator คืออะไร?

```java
// Java — Iterator interface
public class NumberIterator implements Iterator<Integer> {
    public boolean hasNext() { ... }
    public Integer next() { ... }
}

for (int n : iterable) { ... }    // enhanced for-loop
```

```go
// Go 1.23+ — function-based iterator
func Backward[E any](s []E) func(yield func(int, E) bool) {
    return func(yield func(int, E) bool) {
        for i := len(s) - 1; i >= 0; i-- {
            if !yield(i, s[i]) {
                return
            }
        }
    }
}

for i, v := range Backward([]string{"a", "b", "c"}) {
    fmt.Println(i, v)
}
```

## iter.Seq / iter.Seq2 — standard type

```go
import "iter"

// iter.Seq[V]        → yield ค่าเดียว (value only)
func(yield func(V) bool)

// iter.Seq2[K, V]    → yield สองค่า (key + value)
func(yield func(K, V) bool)
```

## สร้าง Custom Iterator

```go
// Iterator แบบ filter
func Filter[E any](s []E, pred func(E) bool) iter.Seq[E] {
    return func(yield func(E) bool) {
        for _, v := range s {
            if pred(v) {
                if !yield(v) {
                    return    // break ออกเมื่อ yield return false
                }
            }
        }
    }
}

for v := range Filter(nums, func(n int) bool { return n > 3 }) {
    fmt.Println(v)
}
```

## ทำไม yield return false ถึงสำคัญ?

เวลาเรียก `break` ใน `for range` → Go จะเรียก `yield` ที่ return `false` → iterator function ต้อง **หยุดทำงาน** เพื่อป้องกัน resource leak

## Iterator กับ Custom Type

```go
type BinaryTree struct {
    Value int
    Left  *BinaryTree
    Right *BinaryTree
}

func (t *BinaryTree) All() iter.Seq[int] {
    return func(yield func(int) bool) {
        t.walk(yield)
    }
}

// ใช้
for v := range tree.All() {
    fmt.Println(v)
}
```

> **Java เทียบ:** `Iterable<T>` + `Iterator<T>` แต่ Go ใช้ function เป็น iterator ไม่ต้องสร้าง class

## ไฟล์ในบทนี้

- `main.go` — Backward, Filter, Enumerate, BinaryTree iterator
