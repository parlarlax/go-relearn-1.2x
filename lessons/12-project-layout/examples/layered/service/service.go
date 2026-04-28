package service

import (
	"fmt"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/model"
	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(name, email string) (*model.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	return s.userRepo.Create(name, email)
}

func (s *UserService) GetByID(id int) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) List() ([]*model.User, error) {
	return s.userRepo.List()
}

func (s *UserService) Update(id int, name, email string) (*model.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	return s.userRepo.Update(id, name, email)
}

func (s *UserService) Delete(id int) error {
	return s.userRepo.Delete(id)
}

type OrderService struct {
	orderRepo repository.OrderRepository
	userRepo  repository.UserRepository
}

func NewOrderService(orderRepo repository.OrderRepository, userRepo repository.UserRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo, userRepo: userRepo}
}

func (s *OrderService) Create(userID int, item string, qty int) (*model.Order, error) {
	if item == "" {
		return nil, fmt.Errorf("item is required")
	}
	if qty <= 0 {
		return nil, fmt.Errorf("qty must be positive")
	}
	if _, err := s.userRepo.GetByID(userID); err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}
	return s.orderRepo.Create(userID, item, qty)
}

func (s *OrderService) ListByUser(userID int) ([]*model.Order, error) {
	return s.orderRepo.ListByUser(userID)
}

func (s *OrderService) List() ([]*model.Order, error) {
	return s.orderRepo.List()
}
