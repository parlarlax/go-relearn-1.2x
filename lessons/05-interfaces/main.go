package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }

func (d Dog) Speak() string { return d.Name + ": Woof!" }

type Cat struct{ Name string }

func (c Cat) Speak() string { return c.Name + ": Meow!" }

type Robot struct{ Model string }

func (r Robot) Speak() string { return r.Model + ": Beep boop!" }

func MakeThemSpeak(speakers ...Speaker) {
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Buffer struct {
	data string
}

func (b *Buffer) Read() string     { return b.data }
func (b *Buffer) Write(d string)   { b.data = d }

func describe(v any) {
	switch v := v.(type) {
	case string:
		fmt.Printf("string: %q (len=%d)\n", v, len(v))
	case int:
		fmt.Printf("int: %d\n", v)
	case bool:
		fmt.Printf("bool: %t\n", v)
	default:
		fmt.Printf("unknown type: %T = %v\n", v, v)
	}
}

type MyError struct{ Msg string }

func (e *MyError) Error() string { return e.Msg }

func nilInterfaceTrap() {
	var err error = nil
	fmt.Println("plain nil:", err == nil)

	var e *MyError = nil
	err = e
	fmt.Println("typed nil:", err == nil)

	if err != nil {
		fmt.Println("TRAP: entered non-nil branch with typed nil!")
	}
}

func main() {
	fmt.Println("=== interface (duck typing) ===")
	d := Dog{"Rex"}
	c := Cat{"Whiskers"}
	r := Robot{"T-800"}
	MakeThemSpeak(d, c, r)

	fmt.Println("\n=== interface composition ===")
	var rw ReadWriter = &Buffer{}
	rw.Write("hello composition")
	fmt.Println(rw.Read())

	fmt.Println("\n=== type switch ===")
	describe("hello")
	describe(42)
	describe(true)
	describe(3.14)

	fmt.Println("\n=== nil interface trap ===")
	nilInterfaceTrap()
}
