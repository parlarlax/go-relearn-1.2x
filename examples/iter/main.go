package main

import (
	"fmt"
	"iter"
)

func Backward[E any](s []E) iter.Seq2[int, E] {
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

func Map[T any, U any](s []T, fn func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for _, v := range s {
			if !yield(fn(v)) {
				return
			}
		}
	}
}

func Take[E any](s []E, n int) iter.Seq[E] {
	return func(yield func(E) bool) {
		for i, v := range s {
			if i >= n {
				return
			}
			if !yield(v) {
				return
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

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		n.walk(yield)
	}
}

func (n *Node) walk(yield func(int) bool) bool {
	if n == nil {
		return true
	}
	if !n.Left.walk(yield) {
		return false
	}
	if !yield(n.Value) {
		return false
	}
	return n.Right.walk(yield)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}

	fmt.Println("=== 1. Backward ===")
	for i, v := range Backward(nums) {
		fmt.Printf("  [%d]=%d\n", i, v)
	}

	fmt.Println("\n=== 2. Filter ===")
	for v := range Filter(nums, func(n int) bool { return n%2 == 0 }) {
		fmt.Printf("  %d\n", v)
	}

	fmt.Println("\n=== 3. Map ===")
	for v := range Map(words, func(s string) int { return len(s) }) {
		fmt.Printf("  len=%d\n", v)
	}

	fmt.Println("\n=== 4. Take (first N) ===")
	for v := range Take(nums, 3) {
		fmt.Printf("  %d\n", v)
	}

	fmt.Println("\n=== 5. Enumerate ===")
	for i, v := range Enumerate(words) {
		fmt.Printf("  #%d: %s\n", i+1, v)
	}

	fmt.Println("\n=== 6. Chaining: Filter → Map ===")
	filtered := []int{}
	for v := range Filter(nums, func(n int) bool { return n%2 == 0 }) {
		filtered = append(filtered, v)
	}
	for v := range Map(filtered, func(n int) int { return n * n }) {
		fmt.Printf("  %d\n", v)
	}

	fmt.Println("\n=== 7. Break early ===")
	count := 0
	for v := range Filter(nums, func(n int) bool { return n > 0 }) {
		fmt.Printf("  %d\n", v)
		count++
		if count >= 3 {
			fmt.Println("  (breaking)")
			break
		}
	}

	fmt.Println("\n=== 8. BinaryTree iterator ===")
	tree := &Node{
		Value: 4,
		Left:  &Node{Value: 2, Left: &Node{Value: 1}, Right: &Node{Value: 3}},
		Right: &Node{Value: 6, Left: &Node{Value: 5}, Right: &Node{Value: 7}},
	}
	for v := range tree.All() {
		fmt.Printf("  %d\n", v)
	}
}
