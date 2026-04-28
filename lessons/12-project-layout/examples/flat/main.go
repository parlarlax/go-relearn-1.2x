package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
	Qty    int    `json:"qty"`
}

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserStore interface {
	Create(name, email string) (*User, error)
	GetByID(id int) (*User, error)
	List() ([]*User, error)
	Update(id int, name, email string) (*User, error)
	Delete(id int) error
}

type OrderStore interface {
	Create(userID int, item string, qty int) (*Order, error)
	ListByUser(userID int) ([]*Order, error)
	List() ([]*Order, error)
}

type MemoryUserStore struct {
	mu    sync.RWMutex
	users map[int]*User
	next  int
}

func NewMemoryUserStore() *MemoryUserStore {
	return &MemoryUserStore{users: make(map[int]*User), next: 1}
}

func (s *MemoryUserStore) Create(name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := &User{ID: s.next, Name: name, Email: email}
	s.users[u.ID] = u
	s.next++
	return u, nil
}

func (s *MemoryUserStore) GetByID(id int) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	return u, nil
}

func (s *MemoryUserStore) List() ([]*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*User, 0, len(s.users))
	for _, u := range s.users {
		result = append(result, u)
	}
	return result, nil
}

func (s *MemoryUserStore) Update(id int, name, email string) (*User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[id]
	if !ok {
		return nil, fmt.Errorf("user %d not found", id)
	}
	u.Name = name
	u.Email = email
	return u, nil
}

func (s *MemoryUserStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[id]
	if !ok {
		return fmt.Errorf("user %d not found", id)
	}
	delete(s.users, id)
	return nil
}

type MemoryOrderStore struct {
	mu     sync.RWMutex
	orders map[int]*Order
	next   int
}

func NewMemoryOrderStore() *MemoryOrderStore {
	return &MemoryOrderStore{orders: make(map[int]*Order), next: 1}
}

func (s *MemoryOrderStore) Create(userID int, item string, qty int) (*Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	o := &Order{ID: s.next, UserID: userID, Item: item, Qty: qty}
	s.orders[o.ID] = o
	s.next++
	return o, nil
}

func (s *MemoryOrderStore) ListByUser(userID int) ([]*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []*Order
	for _, o := range s.orders {
		if o.UserID == userID {
			result = append(result, o)
		}
	}
	return result, nil
}

func (s *MemoryOrderStore) List() ([]*Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]*Order, 0, len(s.orders))
	for _, o := range s.orders {
		result = append(result, o)
	}
	return result, nil
}

type UserService struct {
	store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) Create(name, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required")
	}
	return s.store.Create(name, email)
}

func (s *UserService) GetByID(id int) (*User, error) {
	return s.store.GetByID(id)
}

func (s *UserService) List() ([]*User, error) {
	return s.store.List()
}

func (s *UserService) Update(id int, name, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	return s.store.Update(id, name, email)
}

func (s *UserService) Delete(id int) error {
	return s.store.Delete(id)
}

type OrderService struct {
 orders OrderStore
 users  UserStore
}

func NewOrderService(orders OrderStore, users UserStore) *OrderService {
	return &OrderService{orders: orders, users: users}
}

func (s *OrderService) Create(userID int, item string, qty int) (*Order, error) {
	if item == "" {
		return nil, fmt.Errorf("item is required")
	}
	if qty <= 0 {
		return nil, fmt.Errorf("qty must be positive")
	}
	if _, err := s.users.GetByID(userID); err != nil {
		return nil, fmt.Errorf("user %d not found", userID)
	}
	return s.orders.Create(userID, item, qty)
}

func (s *OrderService) ListByUser(userID int) ([]*Order, error) {
	return s.orders.ListByUser(userID)
}

func (s *OrderService) List() ([]*Order, error) {
	return s.orders.List()
}

func registerUserRoutes(mux *http.ServeMux, svc *UserService) {
	mux.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		users, err := svc.List()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, users)
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		u, err := svc.GetByID(id)
		if err != nil {
			writeErr(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, u)
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		u, err := svc.Create(input.Name, input.Email)
		if err != nil {
			writeErr(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, u)
	})

	mux.HandleFunc("PUT /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		var input struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		u, err := svc.Update(id, input.Name, input.Email)
		if err != nil {
			if errors.Is(err, fmt.Errorf("name is required")) {
				writeErr(w, http.StatusUnprocessableEntity, err.Error())
			} else {
				writeErr(w, http.StatusNotFound, err.Error())
			}
			return
		}
		writeJSON(w, http.StatusOK, u)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := pathID(r)
		if err != nil {
			writeErr(w, http.StatusBadRequest, err.Error())
			return
		}
		if err := svc.Delete(id); err != nil {
			writeErr(w, http.StatusNotFound, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"deleted": "ok"})
	})
}

func registerOrderRoutes(mux *http.ServeMux, svc *OrderService) {
	mux.HandleFunc("GET /orders", func(w http.ResponseWriter, r *http.Request) {
		orders, err := svc.List()
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, orders)
	})

	mux.HandleFunc("GET /users/{id}/orders", func(w http.ResponseWriter, r *http.Request) {
		var userID int
		fmt.Sscanf(r.PathValue("id"), "%d", &userID)
		orders, err := svc.ListByUser(userID)
		if err != nil {
			writeErr(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeJSON(w, http.StatusOK, orders)
	})

	mux.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			UserID int    `json:"user_id"`
			Item   string `json:"item"`
			Qty    int    `json:"qty"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeErr(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		o, err := svc.Create(input.UserID, input.Item, input.Qty)
		if err != nil {
			writeErr(w, http.StatusUnprocessableEntity, err.Error())
			return
		}
		writeJSON(w, http.StatusCreated, o)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(start))
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic recovered", "error", err, "path", r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"status":500,"message":"internal server error"}`))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func chainMiddleware(next http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		next = middlewares[i](next)
	}
	return next
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, &APIError{Status: status, Message: msg})
}

func pathID(r *http.Request) (int, error) {
	var id int
	_, err := fmt.Sscanf(r.PathValue("id"), "%d", &id)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid id")
	}
	return id, nil
}

func main() {
	port, _ := strconv.Atoi(getEnv("PORT", "8081"))
	addr := fmt.Sprintf(":%d", port)
	slog.Info("starting flat server", "addr", addr)

	var userStore UserStore = NewMemoryUserStore()
	var orderStore OrderStore = NewMemoryOrderStore()

	userSvc := NewUserService(userStore)
	orderSvc := NewOrderService(orderStore, userStore)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})
	registerUserRoutes(mux, userSvc)
	registerOrderRoutes(mux, orderSvc)

	handler := chainMiddleware(mux, loggingMiddleware, recoveryMiddleware)

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	slog.Info("received signal", "signal", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "error", err)
	}
	slog.Info("server stopped")
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
