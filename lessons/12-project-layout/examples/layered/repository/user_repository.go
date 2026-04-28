package repository

import "sync"

type UserRepository struct {
	mu    sync.RWMutex
	users map[int]*User
	next  int
}

type User = struct {
	ID    int
	Name  string
	Email string
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[int]*User),
		next:  1,
	}
}

func (r *UserRepository) Create(name, email string) *User {
	r.mu.Lock()
	defer r.mu.Unlock()
	u := &User{ID: r.next, Name: name, Email: email}
	r.users[u.ID] = u
	r.next++
	return u
}

func (r *UserRepository) GetByID(id int) (*User, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id]
	return u, ok
}

func (r *UserRepository) GetAll() []*User {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make([]*User, 0, len(r.users))
	for _, u := range r.users {
		result = append(result, u)
	}
	return result
}

func (r *UserRepository) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, ok := r.users[id]
	if ok {
		delete(r.users, id)
	}
	return ok
}
