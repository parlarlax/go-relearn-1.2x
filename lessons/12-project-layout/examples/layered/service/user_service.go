package service

import (
	"fmt"

	rep "github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/repository"
)

type UserService struct {
	repo *rep.UserRepository
}

func NewUserService(repo *rep.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(name, email string) (*rep.User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	return s.repo.Create(name, email), nil
}

func (s *UserService) GetByID(id int) (*rep.User, error) {
	u, ok := s.repo.GetByID(id)
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (s *UserService) List() []*rep.User {
	return s.repo.GetAll()
}

func (s *UserService) Delete(id int) error {
	if !s.repo.Delete(id) {
		return fmt.Errorf("user %d not found", id)
	}
	return nil
}
