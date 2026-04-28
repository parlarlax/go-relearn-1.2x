package main

import (
	"errors"
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func namedReturn(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func deferExample() {
	fmt.Println("step 1")
	defer fmt.Println("deferred: runs last")
	defer fmt.Println("deferred: runs second-to-last (LIFO)")
	fmt.Println("step 2")
}

func fileReadExample() {
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		fmt.Println("expected error:", err)
		return
	}
	defer file.Close()
}

func init() {
	fmt.Println("init() called — runs before main()")
}

func main() {
	fmt.Println("=== basic function ===")
	fmt.Println("add(3, 4) =", add(3, 4))

	fmt.Println("\n=== multiple return ===")
	result, err := divide(10, 3)
	fmt.Printf("10 / 3 = %.2f, err = %v\n", result, err)

	result, err = divide(10, 0)
	fmt.Printf("10 / 0 = %.2f, err = %v\n", result, err)

	fmt.Println("\n=== named return ===")
	x, y := namedReturn(1, 2)
	fmt.Println("namedReturn(1, 2) =", x, y)

	fmt.Println("\n=== defer ===")
	deferExample()

	fmt.Println("\n=== file + defer ===")
	fileReadExample()
}
