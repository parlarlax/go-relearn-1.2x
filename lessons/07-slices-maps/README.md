# บทที่ 7: Arrays, Slices, Maps

> Go ไม่มี generics collections แบบ Java (`ArrayList<E>`, `HashMap<K,V>`) — ใช้ slice และ map ในภาษาเลย (built-in)

## Array vs Slice

```java
// Java — ใช้ ArrayList เกือบตลอด
String[] arr = new String[3];           // fixed array
List<String> list = new ArrayList<>();  // dynamic
list.add("hello");
list.get(0);
```

```go
// Go
arr := [3]string{"a", "b", "c"}   // Array — ความยาวตายตัว (ใช้น้อย)
slice := []string{"a", "b", "c"}  // Slice — dynamic (ใช้เกือบตลอด)
slice = append(slice, "d")        // เพิ่ม element
fmt.Println(slice[0])             // เข้าถึง
```

> **Array = ความยาวตายตัว, Slice = ความยาวเปลี่ยนได้** — ใช้ slice 99% ของเวลา

## Slice — เทียบ ArrayList

```go
nums := []int{1, 2, 3, 4, 5}

len(nums)                        // 5  — เหมือน list.size()
nums[1:3]                        // [2, 3] — slice (เหมือน subList)
nums[:3]                         // [1, 2, 3] — ตั้งแต่ต้นถึง index 3
nums[2:]                         // [3, 4, 5] — ตั้งแต่ index 2 ถึงท้าย

nums = append(nums, 6)           // เพิ่มตัวเดียว
nums = append(nums, 7, 8, 9)    // เพิ่มหลายตัว

other := []int{10, 11}
nums = append(nums, other...)    // เพิ่มทั้ง slice (เหมือน list.addAll())
```

## Slice Internals — สำคัญ!

Slice ใน Go มี 3 ส่วน:
- **pointer** → ชี้ไปที่ underlying array
- **length** (`len`) → จำนวน element ตอนนี้
- **capacity** (`cap`) → ความจุที่ allocate ไว้

```go
s := make([]int, 0, 10)  // length=0, capacity=10
```

> **Java เทียบ:** `new ArrayList<>(10)` — กำหนด initial capacity

### ระวัง: Slice แชร์ underlying array!

```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]     // b = [2, 3]
b[0] = 99       // a เปลี่ยนด้วย! a = [1, 99, 3, 4, 5]

// แก้: ใช้ copy
c := make([]int, len(b))
copy(c, b)
```

## Map — เทียบ HashMap

```java
// Java
Map<String, Integer> scores = new HashMap<>();
scores.put("Alice", 100);
scores.get("Alice");          // 100 หรือ null
scores.containsKey("Alice");  // true
```

```go
// Go
scores := map[string]int{        // map[keyType]valueType
    "Alice": 100,
    "Bob":   85,
}

scores["Charlie"] = 90           // put
val := scores["Alice"]           // get (ถ้าไม่มี → zero value ของ value type)
val, ok := scores["David"]       // get + check existence
if !ok {
    fmt.Println("not found")     // ok=false ถ้าไม่มี key
}

delete(scores, "Bob")            // remove

for key, val := range scores {   // iterate
    fmt.Println(key, val)
}
```

> **สำคัญ:** `scores["missing"]` ไม่ panic — ได้ zero value ของ value type กลับมา

## Range (for-each)

```java
// Java
for (String item : list) { ... }
for (Map.Entry<String, Integer> e : map.entrySet()) { ... }
```

```go
// Go
for index, value := range slice { ... }
for key, value := range myMap  { ... }
for _, value := range slice    { ... }  // ไม่สน index
for index, _ := range slice    { ... }  // ไม่สน value
```

## ไฟล์ในบทนี้

- `main.go` — slice, map, range, copy, make
