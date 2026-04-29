package application

import (
	"fmt"

	"github.com/lax/go-relearn/examples/gontainer-ddd/domain"
)

type OrderService struct {
	orders domain.OrderStore
	users  domain.UserStore
}

func NewOrderService(orders domain.OrderStore, users domain.UserStore) *OrderService {
	return &OrderService{orders: orders, users: users}
}

func (s *OrderService) Create(userID int, item string, qty int) (*domain.Order, error) {
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

func (s *OrderService) List() ([]*domain.Order, error) {
	return s.orders.List()
}

func (s *OrderService) ListByUser(userID int) ([]*domain.Order, error) {
	return s.orders.ListByUser(userID)
}
