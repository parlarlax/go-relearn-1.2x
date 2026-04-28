package generics

import (
	"cmp"
	"fmt"
)

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func ExampleConstraints() {
	fmt.Println("min:", Min(3, 7))
	fmt.Println("min:", Min("apple", "banana"))
	fmt.Println("max:", Max(3.14, 2.71))

	fmt.Println("contains:", Contains([]string{"a", "b", "c"}, "b"))
	fmt.Println("contains:", Contains([]int{1, 2, 3}, 5))
}
