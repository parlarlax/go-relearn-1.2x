package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func birthday(u *User) {
	u.Age++
}

func tryModify(val int) {
	val = 999
}

func modifyByPointer(ptr *int) {
	*ptr = 999
}

func main() {
	fmt.Println("=== pointer basics ===")
	x := 42
	p := &x
	fmt.Println("x:", x, "p:", p, "*p:", *p)
	*p = 100
	fmt.Println("after *p = 100, x:", x)

	fmt.Println("\n=== struct: value vs pointer ===")
	a := User{Name: "Alice", Age: 25}
	b := a
	b.Name = "Bob"
	fmt.Println("a (original):", a.Name)
	fmt.Println("b (copy):", b.Name)

	c := &User{Name: "Alice", Age: 25}
	d := c
	d.Name = "Charlie"
	fmt.Println("c (pointer):", c.Name)
	fmt.Println("d (same ref):", d.Name)

	fmt.Println("\n=== function: pass by value vs pointer ===")
	val := 10
	tryModify(val)
	fmt.Println("after tryModify:", val)
	modifyByPointer(&val)
	fmt.Println("after modifyByPointer:", val)

	fmt.Println("\n=== pointer receiver pattern ===")
	user := User{Name: "Dave", Age: 30}
	fmt.Println("before:", user.Age)
	birthday(&user)
	fmt.Println("after birthday:", user.Age)

	fmt.Println("\n=== nil pointer ===")
	var ptr *User
	fmt.Println("nil pointer:", ptr)
	if ptr != nil {
		fmt.Println(ptr.Name)
	} else {
		fmt.Println("pointer is nil — safe to check")
	}
}
