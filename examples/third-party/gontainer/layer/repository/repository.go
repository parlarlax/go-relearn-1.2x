package repository

import (
	"fmt"
	"sync"

	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/model"
)

type UserRepository struct {
	mu    sync.RWMutex
	users map[int]*model.User
	next  int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[int]*model.User), next: 1}
}

func (r *UserRepository) Create(name, email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u := &model.User{ID: r.next, Name: name, Email: email}
	r.users[u.ID] = u
	r.next++
	return u, nil
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (r *UserRepository) List() ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*model.User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}
	return result, nil
}

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[int]*model.Order
	next   int
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{orders: make(map[int]*model.Order), next: 1}
}

func (r *OrderRepository) Create(userID int, item string, qty int) (*model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	o := &model.Order{ID: r.next, UserID: userID, Item: item, Qty: qty}
	r.orders[o.ID] = o
	r.next++
	return o, nil
}

func (r *OrderRepository) ListByUser(userID int) ([]*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*model.Order
	for _, o := range r.orders {
		if o.UserID == userID {
			result = append(result, o)
		}
	}
	return result, nil
}

func (r *OrderRepository) List() ([]*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*model.Order, 0, len(r.orders))
	for _, o := range r.orders {
		result = append(result, o)
	}
	return result, nil
}
