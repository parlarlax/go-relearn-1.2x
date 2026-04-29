package service

import (
	"fmt"

	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/model"
	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(name, email string) (*model.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	return s.repo.Create(name, email)
}

func (s *UserService) Get(id int) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) List() ([]*model.User, error) {
	return s.repo.List()
}

type OrderService struct {
	orderRepo *repository.OrderRepository
	userRepo  *repository.UserRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, userRepo *repository.UserRepository) *OrderService {
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

func (s *OrderService) List() ([]*model.Order, error) {
	return s.orderRepo.List()
}

func (s *OrderService) ListByUser(userID int) ([]*model.Order, error) {
	return s.orderRepo.ListByUser(userID)
}
