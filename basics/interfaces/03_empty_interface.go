package interfaces

import "fmt"

func PrintAny(v interface{}) {
	fmt.Printf("value: %v, type: %T\n", v, v)
}

func ExampleEmptyInterface() {
	PrintAny(42)
	PrintAny("hello")
	PrintAny([]int{1, 2, 3})
	PrintAny(struct{ Name string }{"test"})
}
