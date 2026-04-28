package main

import (
	"fmt"
)

func Backward[E any](s []E) func(yield func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func Filter[E any](s []E, predicate func(E) bool) func(yield func(E) bool) {
	return func(yield func(E) bool) {
		for _, v := range s {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Enumerate[E any](s []E) func(yield func(int, E) bool) {
	return func(yield func(int, E) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	fruits := []string{"apple", "banana", "cherry", "date"}

	fmt.Println("=== Backward ===")
	for i, v := range Backward(fruits) {
		fmt.Printf("  [%d] %s\n", i, v)
	}

	fmt.Println("\n=== Filter (length > 5) ===")
	for v := range Filter(fruits, func(s string) bool { return len(s) > 5 }) {
		fmt.Printf("  %s\n", v)
	}

	fmt.Println("\n=== Enumerate ===")
	for i, v := range Enumerate(fruits) {
		fmt.Printf("  #%d: %s\n", i+1, v)
	}
}
