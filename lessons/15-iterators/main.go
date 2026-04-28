package main

import (
	"fmt"
	"iter"
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

func Filter[E any](s []E, pred func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range s {
			if pred(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func Enumerate[E any](s []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (t *BinaryTree) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		t.walk(yield)
	}
}

func (t *BinaryTree) walk(yield func(int) bool) bool {
	if t == nil {
		return true
	}
	if !t.Left.walk(yield) {
		return false
	}
	if !yield(t.Value) {
		return false
	}
	return t.Right.walk(yield)
}

func main() {
	fruits := []string{"apple", "banana", "cherry", "date", "elderberry"}

	fmt.Println("=== Backward ===")
	for i, v := range Backward(fruits) {
		fmt.Printf("  [%d] %s\n", i, v)
	}

	fmt.Println("\n=== Filter (len > 5) ===")
	for v := range Filter(fruits, func(s string) bool { return len(s) > 5 }) {
		fmt.Printf("  %s\n", v)
	}

	fmt.Println("\n=== Enumerate ===")
	for i, v := range Enumerate(fruits) {
		fmt.Printf("  #%d: %s\n", i+1, v)
	}

	fmt.Println("\n=== Break early (yield returns false) ===")
	for i, v := range Enumerate(fruits) {
		fmt.Printf("  %s\n", v)
		if v == "cherry" {
			fmt.Printf("  (breaking at index %d)\n", i)
			break
		}
	}

	fmt.Println("\n=== BinaryTree iterator ===")
	tree := &BinaryTree{
		Value: 4,
		Left: &BinaryTree{
			Value: 2,
			Left:  &BinaryTree{Value: 1},
			Right: &BinaryTree{Value: 3},
		},
		Right: &BinaryTree{
			Value: 6,
			Left:  &BinaryTree{Value: 5},
			Right: &BinaryTree{Value: 7},
		},
	}
	for v := range tree.All() {
		fmt.Printf("  %d\n", v)
	}
}
