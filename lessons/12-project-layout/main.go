package main

import (
	"fmt"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(id int, name, email string) *User {
	return &User{ID: id, Name: name, Email: email}
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

type UserService struct {
	users map[int]*User
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}

func (s *UserService) Create(u *User) error {
	if err := u.Validate(); err != nil {
		return err
	}
	s.users[u.ID] = u
	return nil
}

func (s *UserService) GetByID(id int) (*User, error) {
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (s *UserService) List() []*User {
	result := make([]*User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result
}

func main() {
	service := NewUserService()

	service.Create(NewUser(1, "Alice", "alice@test.com"))
	service.Create(NewUser(2, "Bob", "bob@test.com"))
	service.Create(NewUser(3, "Charlie", "charlie@test.com"))

	if user, err := service.GetByID(2); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("found: %+v\n", user)
	}

	fmt.Println("\nall users:")
	for _, u := range service.List() {
		fmt.Printf("  [%d] %s (%s)\n", u.ID, u.Name, u.Email)
	}

	if _, err := service.GetByID(99); err != nil {
		fmt.Println("\nexpected error:", err)
	}
}
