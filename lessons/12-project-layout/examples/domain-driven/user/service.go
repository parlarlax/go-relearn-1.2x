package user

import (
	"fmt"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/store"
)

type Service struct {
	users store.UserStore
}

func NewService(users store.UserStore) *Service {
	return &Service{users: users}
}

func (s *Service) Create(name, email string) (*store.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	return s.users.Create(name, email)
}

func (s *Service) GetByID(id int) (*store.User, error) {
	return s.users.GetByID(id)
}

func (s *Service) List() ([]*store.User, error) {
	return s.users.List()
}

func (s *Service) Update(id int, name, email string) (*store.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	return s.users.Update(id, name, email)
}

func (s *Service) Delete(id int) error {
	return s.users.Delete(id)
}
