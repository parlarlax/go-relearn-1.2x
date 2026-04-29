package application

import (
	"fmt"

	"github.com/lax/go-relearn/examples/third-party/gontainer/ddd/domain"
)

type UserService struct {
	users domain.UserStore
}

func NewUserService(users domain.UserStore) *UserService {
	return &UserService{users: users}
}

func (s *UserService) Create(name, email string) (*domain.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	return s.users.Create(name, email)
}

func (s *UserService) Get(id int) (*domain.User, error) {
	return s.users.GetByID(id)
}

func (s *UserService) List() ([]*domain.User, error) {
	return s.users.List()
}
