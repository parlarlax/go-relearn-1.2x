# บทที่ 8: Pointers

> Java ซ่อน pointer ไว้ — Go ให้คุณใช้ pointer ได้โดยตรง แต่ **ไม่มี pointer arithmetic** (ปลอดภัยกว่า C/C++)

## Java Object vs Go Pointer

```java
// Java — ทุก object เป็น reference อัตโนมัติ
User a = new User("Alice");
User b = a;        // b ชี้ไปที่ object เดียวกัน
b.name = "Bob";
System.out.println(a.name);  // "Bob" — เปลี่ยนเดียวกัน
```

```go
// Go — ต้องระบุ pointer เองด้วย &
a := User{Name: "Alice"}
b := a                   // copy ค่า (ไม่ใช่ reference!)
b.Name = "Bob"
fmt.Println(a.Name)      // "Alice" — a ไม่เปลี่ยน

// ใช้ pointer
c := &User{Name: "Alice"}   // *User
d := c                      // d ชี้ไปที่เดียวกัน
d.Name = "Charlie"
fmt.Println(c.Name)          // "Charlie" — เปลี่ยนเดียวกัน
```

## Syntax

```go
x := 42
p := &x     // & = "address of" → p เป็น *int (pointer to int)
*p = 100    // * = "dereference" → แก้ค่าที่ p ชี้อยู่
fmt.Println(x)   // 100
```

| Syntax | ความหมาย | Java เทียบ |
|---|---|---|
| `&x` | address of x | (Java ซ่อนไว้) |
| `*p` | value ที่ pointer ชี้ไป | dereference reference |
| `*Type` | pointer type | reference type |
| `nil` | null pointer | `null` |

## เมื่อไหร่ต้องใช้ pointer?

1. **อยากให้ function แก้ค่าตัวแปรเดิมได้**
2. **หลีกเลี่ยง copy struct ใหญ่** (performance)
3. **ต้องการ nil เป็นค่า default** (zero value ของ pointer = nil)

## Pointer กับ Struct

```go
func birthday(u *User) {
    u.Age++   // Go อำนวยความสะดวก → (*u).Age เขียนเป็น u.Age ได้
}

user := User{Name: "Alice", Age: 25}
birthday(&user)
fmt.Println(user.Age)   // 26
```

> Go **ไม่มี pointer arithmetic** (เช่น `p++` หรือ `p+1`) ต่างจาก C — นี่ทำให้ปลอดภัยกว่า

## ไฟล์ในบทนี้

- `main.go` — pointer basics, pointer with struct, function with pointer
