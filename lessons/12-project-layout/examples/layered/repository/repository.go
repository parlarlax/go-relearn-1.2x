package repository

import (
	"fmt"
	"sync"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/model"
)

type UserRepository interface {
	Create(name, email string) (*model.User, error)
	GetByID(id int) (*model.User, error)
	List() ([]*model.User, error)
	Update(id int, name, email string) (*model.User, error)
	Delete(id int) error
}

type OrderRepository interface {
	Create(userID int, item string, qty int) (*model.Order, error)
	ListByUser(userID int) ([]*model.Order, error)
	List() ([]*model.Order, error)
}

type MemoryUserRepository struct {
	mu    sync.RWMutex
	users map[int]*model.User
	next  int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[int]*model.User),
		next:  1,
	}
}

func (r *MemoryUserRepository) Create(name, email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u := &model.User{ID: r.next, Name: name, Email: email}
	r.users[u.ID] = u
	r.next++
	return u, nil
}

func (r *MemoryUserRepository) GetByID(id int) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (r *MemoryUserRepository) List() ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*model.User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}
	return result, nil
}

func (r *MemoryUserRepository) Update(id int, name, email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	u, ok := r.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	u.Name = name
	u.Email = email
	return u, nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.users[id]
	if !ok {
		return fmt.Errorf("user %d not found", id)
	}
	delete(r.users, id)
	return nil
}

type MemoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[int]*model.Order
	next   int
}

func NewMemoryOrderRepository() *MemoryOrderRepository {
	return &MemoryOrderRepository{
		orders: make(map[int]*model.Order),
		next:   1,
	}
}

func (r *MemoryOrderRepository) Create(userID int, item string, qty int) (*model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	o := &model.Order{ID: r.next, UserID: userID, Item: item, Qty: qty}
	r.orders[o.ID] = o
	r.next++
	return o, nil
}

func (r *MemoryOrderRepository) ListByUser(userID int) ([]*model.Order, error) {
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

func (r *MemoryOrderRepository) List() ([]*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*model.Order, 0, len(r.orders))
	for _, o := range r.orders {
		result = append(result, o)
	}
	return result, nil
}
