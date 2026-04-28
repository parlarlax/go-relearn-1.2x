package db

import "sync"

type Store[T any] struct {
	mu    sync.RWMutex
	items map[int]T
	next  int
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{
		items: make(map[int]T),
		next:  1,
	}
}

func (s *Store[T]) Create(item T) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.next
	s.items[id] = item
	s.next++
	return id
}

func (s *Store[T]) Get(id int) (T, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.items[id]
	return item, ok
}

func (s *Store[T]) List() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]T, 0, len(s.items))
	for _, item := range s.items {
		result = append(result, item)
	}
	return result
}

func (s *Store[T]) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[id]
	if ok {
		delete(s.items, id)
	}
	return ok
}
