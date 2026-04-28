package generics

import "fmt"

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type OrderedMap[K comparable, V any] struct {
	pairs []Pair[K, V]
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{}
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	for i, p := range m.pairs {
		if p.Key == key {
			m.pairs[i].Value = value
			return
		}
	}
	m.pairs = append(m.pairs, Pair[K, V]{Key: key, Value: value})
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	for _, p := range m.pairs {
		if p.Key == key {
			return p.Value, true
		}
	}
	var zero V
	return zero, false
}

func (m *OrderedMap[K, V]) Keys() []K {
	keys := make([]K, len(m.pairs))
	for i, p := range m.pairs {
		keys[i] = p.Key
	}
	return keys
}

func ExampleOrderedMap() {
	m := NewOrderedMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)

	if v, ok := m.Get("two"); ok {
		fmt.Println("two:", v)
	}

	fmt.Println("keys:", m.Keys())
}
