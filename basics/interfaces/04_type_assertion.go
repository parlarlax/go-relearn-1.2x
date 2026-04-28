package interfaces

import "fmt"

func Describe(i interface{}) string {
	switch v := i.(type) {
	case string:
		return fmt.Sprintf("string: %q (len=%d)", v, len(v))
	case int:
		return fmt.Sprintf("int: %d", v)
	case bool:
		return fmt.Sprintf("bool: %t", v)
	default:
		return fmt.Sprintf("unknown: %T", v)
	}
}

func ExampleTypeAssertion() {
	fmt.Println(Describe("hello"))
	fmt.Println(Describe(42))
	fmt.Println(Describe(true))
	fmt.Println(Describe(3.14))
}
