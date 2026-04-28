package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Store struct {
	mu    sync.RWMutex
	users map[int]User
	next  int
}

func NewStore() *Store {
	return &Store{
		users: make(map[int]User),
		next:  1,
	}
}

func (s *Store) Create(name string) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := User{ID: s.next, Name: name}
	s.users[u.ID] = u
	s.next++
	return u
}

func (s *Store) Get(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *Store) List() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[id]
	if ok {
		delete(s.users, id)
	}
	return ok
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}

func main() {
	store := NewStore()
	store.Create("Alice")
	store.Create("Bob")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"message": "Go HTTP Server running!",
		})
	})

	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, store.List())
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id int
		fmt.Sscanf(r.PathValue("id"), "%d", &id)
		user, ok := store.Get(id)
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, http.StatusOK, user)
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"name"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Name == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name required"})
			return
		}
		user := store.Create(input.Name)
		writeJSON(w, http.StatusCreated, user)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		var id int
		fmt.Sscanf(r.PathValue("id"), "%d", &id)
		if !store.Delete(id) {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"deleted": "ok"})
	})

	handler := loggingMiddleware(mux)

	slog.Info("server starting on :8080")
	fmt.Println("Try:")
	fmt.Println("  curl http://localhost:8080/users")
	fmt.Println("  curl http://localhost:8080/users/1")
	fmt.Println(`  curl -X POST -d '{"name":"Charlie"}' http://localhost:8080/users`)
	fmt.Println("  curl -X DELETE http://localhost:8080/users/1")

	http.ListenAndServe(":8080", handler)
}
