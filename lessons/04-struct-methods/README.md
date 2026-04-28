# บทที่ 4: Struct & Methods

> Go **ไม่มี class** — ใช้ `struct` + `method` แทน แต่ concept คล้ายกันมาก

## Struct = Java Class (แบบไม่มี inheritance)

```java
// Java
public class User {
    private String name;
    private int age;
    
    public User(String name, int age) {
        this.name = name;
        this.age = age;
    }
    
    public String greet() {
        return "Hi, I'm " + name;
    }
}
```

```go
// Go
type User struct {
    Name string  // Go ไม่มี private/public keyword — ใช้ตัวพิมพ์ใหญ่/เล็กแทน
    Age  int     // ตัวพิมพ์ใหญ่ = public (exported), ตัวเล็ก = private (unexported)
}

// Constructor — Go ไม่มี constructor ต้องสร้างเอง
func NewUser(name string, age int) *User {
    return &User{Name: name, Age: age}
}

// Method — ผูกกับ struct ด้วย "receiver"
func (u User) Greet() string {
    return fmt.Sprintf("Hi, I'm %s", u.Name)
}
```

## Visibility — ตัวพิมพ์ใหญ่ vs เล็ก

```go
type User struct {
    Name  string  // exported (public) — ตัวอักษรตัวแรกพิมพ์ใหญ่
    email string  // unexported (private) — ตัวอักษรตัวแรกพิมพ์เล็ก
}
```

> นี่คือสิ่งที่ Java ใช้ `public` / `private` keyword แต่ Go ใช้ **อักษรตัวแรก** แทน — ใช้ได้กับทุกอย่าง (struct, function, method, field, constant)

## Value Receiver vs Pointer Receiver

```go
// Value receiver — ส่งสำเนา (เหมือน pass-by-value ใน Java primitive)
func (u User) Greet() string {
    return u.Name
}

// Pointer receiver — ส่ง reference (เหมือน pass-by-reference ของ Java object)
func (u *User) Birthday() {
    u.Age++  // แก้ค่าจริงได้
}
```

**กฎง่ายๆ:**
- อยาก **แก้ข้อมูล** → pointer receiver `*User`
- อยาก **หลีกเลี่ยง copy ข้อมูลใหญ่** → pointer receiver
- อยาก **แค่อ่าน** และ struct เล็ก → value receiver ได้

> **Java เทียบ:** ทุก object ใน Java ส่งแบบ reference อยู่แล้ว ส่วน primitive ส่งแบบ value — Go ให้เลือกเองทุก type

## Struct Embedding (Composition แทน Inheritance)

```java
// Java — ใช้ extends
public class Admin extends User {
    private String role;
}
```

```go
// Go — ใช้ embedding (composition)
type Admin struct {
    User        // ← embed User เข้าไป (ไม่ใช่ inheritance)
    Role string
}

admin := Admin{
    User: User{Name: "Boss", Age: 40},
    Role: "superadmin",
}
fmt.Println(admin.Name)  // เข้าถึง field ของ User ได้ตรงๆ
admin.Greet()            // เรียก method ของ User ได้ด้วย
```

> Go ไม่มี `extends` — ใช้ **composition** แทน inheritance ทั้งหมด

## ไฟล์ในบทนี้

- `main.go` — struct, methods, value/pointer receiver, embedding
