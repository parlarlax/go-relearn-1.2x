package main

import (
	"fmt"
	"iter"
)

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

	fmt.Println("=== In-order traversal via iterator ===")
	for v := range tree.All() {
		fmt.Printf("  %d\n", v)
	}
}
