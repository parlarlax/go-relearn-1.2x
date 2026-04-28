package main

import (
	"errors"
	"fmt"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func findUser(id int) (string, error) {
	if id <= 0 {
		return "", &AppError{Code: 400, Message: "invalid id"}
	}
	if id > 100 {
		return "", fmt.Errorf("user %d not found", id)
	}
	return fmt.Sprintf("User_%d", id), nil
}

func wrapExample() error {
	inner := errors.New("connection refused")
	return fmt.Errorf("db query failed: %w", inner)
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()
	result = a / b
	return result, nil
}

func main() {
	fmt.Println("=== basic error handling ===")
	if name, err := findUser(1); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("found:", name)
	}

	if _, err := findUser(-1); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\n=== errors.As — extract custom error ===")
	_, err := findUser(-1)
	var appErr *AppError
	if errors.As(err, &appErr) {
		fmt.Printf("app error: code=%d msg=%s\n", appErr.Code, appErr.Message)
	}

	fmt.Println("\n=== errors.Is + wrap/unwrap ===")
	wrapped := wrapExample()
	fmt.Println("wrapped:", wrapped)
	fmt.Println("unwrapped:", errors.Unwrap(wrapped))

	original := errors.New("connection refused")
	fmt.Println("is original?", errors.Is(wrapped, original))

	fmt.Println("\n=== panic/recover ===")
	if result, err := safeDivide(10, 2); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10/2 =", result)
	}

	if result, err := safeDivide(10, 0); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10/0 =", result)
	}
}
