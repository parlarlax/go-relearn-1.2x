package main

import "errors"

func Add(a, b int) int {
	return a + b
}

var ErrDivisionByZero = errors.New("division by zero")

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}
