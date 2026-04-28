package interfaces

import "fmt"

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return e.Msg
}

func ExampleNilInterface() {
	var err error = nil
	fmt.Println("nil interface:", err)

	var e *MyError = nil
	err = e
	fmt.Println("typed nil:", err)
	fmt.Println("is nil?", err == nil)

	if err != nil {
		fmt.Println("entered non-nil branch with typed nil!")
	}
}
