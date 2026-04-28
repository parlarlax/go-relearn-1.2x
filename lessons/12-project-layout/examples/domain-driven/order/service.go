package order

import (
	"fmt"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/store"
)

type Service struct {
	orders store.OrderStore
	users  store.UserStore
}

func NewService(orders store.OrderStore, users store.UserStore) *Service {
	return &Service{orders: orders, users: users}
}

func (s *Service) Create(userID int, item string, qty int) (*store.Order, error) {
	if item == "" {
		return nil, fmt.Errorf("item is required")
	}
	if qty <= 0 {
		return nil, fmt.Errorf("qty must be positive")
	}
	if _, err := s.users.GetByID(userID); err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}
	return s.orders.Create(userID, item, qty)
}

func (s *Service) ListByUser(userID int) ([]*store.Order, error) {
	return s.orders.ListByUser(userID)
}

func (s *Service) List() ([]*store.Order, error) {
	return s.orders.List()
}
