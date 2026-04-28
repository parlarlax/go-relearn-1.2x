package store

import (
	"fmt"
	"sync"
)

type MemoryUserStore struct {
	mu    sync.RWMutex
	users map[int]*User
	next  int
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{
		users: make(map[int]*User),
		next:  1,
	}
}

func (s *MemoryUserStore) Create(name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := &User{ID: s.next, Name: name, Email: email}
	s.users[u.ID] = u
	s.next++
	return u, nil
}

func (s *MemoryUserStore) GetByID(id int) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (s *MemoryUserStore) List() ([]*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result, nil
}

func (s *MemoryUserStore) Update(id int, name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	u.Name = name
	u.Email = email
	return u, nil
}

func (s *MemoryUserStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[id]
	if !ok {
		return fmt.Errorf("user %d not found", id)
	}
	delete(s.users, id)
	return nil
}

type MemoryOrderStore struct {
	mu     sync.RWMutex
	orders map[int]*Order
	next   int
}

func NewMemoryOrderStore() *MemoryOrderStore {
	return &MemoryOrderStore{
		orders: make(map[int]*Order),
		next:   1,
	}
}

func (s *MemoryOrderStore) Create(userID int, item string, qty int) (*Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o := &Order{ID: s.next, UserID: userID, Item: item, Qty: qty}
	s.orders[o.ID] = o
	s.next++
	return o, nil
}

func (s *MemoryOrderStore) ListByUser(userID int) ([]*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []*Order
	for _, o := range s.orders {
		if o.UserID == userID {
			result = append(result, o)
		}
	}
	return result, nil
}

func (s *MemoryOrderStore) List() ([]*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*Order, 0, len(s.orders))
	for _, o := range s.orders {
		result = append(result, o)
	}
	return result, nil
}
