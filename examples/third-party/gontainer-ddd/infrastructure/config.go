package infrastructure

import (
	"fmt"

	"github.com/lax/go-relearn/examples/third-party/gontainer-ddd/domain"
)

type Config struct {
	Port int
}

func NewConfig() *Config {
	return &Config{Port: 18080}
}

func (c *Config) Addr() string {
	return fmt.Sprintf(":%d", c.Port)
}

type MemoryUserStore struct {
	users map[int]*domain.User
	next  int
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{users: make(map[int]*domain.User), next: 1}
}

func (s *MemoryUserStore) Create(name, email string) (*domain.User, error) {
	u := &domain.User{ID: s.next, Name: name, Email: email}
	s.users[u.ID] = u
	s.next++
	return u, nil
}

func (s *MemoryUserStore) GetByID(id int) (*domain.User, error) {
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (s *MemoryUserStore) List() ([]*domain.User, error) {
	result := make([]*domain.User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result, nil
}

type MemoryOrderStore struct {
	orders map[int]*domain.Order
	next   int
}

func NewMemoryOrderStore() *MemoryOrderStore {
	return &MemoryOrderStore{orders: make(map[int]*domain.Order), next: 1}
}

func (s *MemoryOrderStore) Create(userID int, item string, qty int) (*domain.Order, error) {
	o := &domain.Order{ID: s.next, UserID: userID, Item: item, Qty: qty}
	s.orders[o.ID] = o
	s.next++
	return o, nil
}

func (s *MemoryOrderStore) ListByUser(userID int) ([]*domain.Order, error) {
	var result []*domain.Order
	for _, o := range s.orders {
		if o.UserID == userID {
			result = append(result, o)
		}
	}
	return result, nil
}

func (s *MemoryOrderStore) List() ([]*domain.Order, error) {
	result := make([]*domain.Order, 0, len(s.orders))
	for _, o := range s.orders {
		result = append(result, o)
	}
	return result, nil
}
