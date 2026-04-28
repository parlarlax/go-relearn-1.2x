package generics

import "fmt"

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func ExampleStack() {
	s := NewStack[int]()
	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("len:", s.Len())

	if v, ok := s.Peek(); ok {
		fmt.Println("peek:", v)
	}

	if v, ok := s.Pop(); ok {
		fmt.Println("pop:", v)
	}

	fmt.Println("len after pop:", s.Len())
}
