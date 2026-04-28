package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	users   = make(map[int]*User)
	usersMu sync.RWMutex
	nextID  = 1
)

func jsonResp(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func jsonErr(w http.ResponseWriter, status int, msg string) {
	jsonResp(w, status, map[string]string{"error": msg})
}

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	usersMu.RLock()
	defer usersMu.RUnlock()
	list := make([]*User, 0, len(users))
	for _, u := range users {
		list = append(list, u)
	}
	jsonResp(w, http.StatusOK, list)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Sscanf(r.PathValue("id"), "%d", &id)
	usersMu.RLock()
	u, ok := users[id]
	usersMu.RUnlock()
	if !ok {
		jsonErr(w, http.StatusNotFound, "user not found")
		return
	}
	jsonResp(w, http.StatusOK, u)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Name == "" {
		jsonErr(w, http.StatusBadRequest, "name and email required")
		return
	}
	usersMu.Lock()
	u := &User{ID: nextID, Name: input.Name, Email: input.Email}
	users[u.ID] = u
	nextID++
	usersMu.Unlock()
	jsonResp(w, http.StatusCreated, u)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Sscanf(r.PathValue("id"), "%d", &id)
	usersMu.Lock()
	_, ok := users[id]
	if ok {
		delete(users, id)
	}
	usersMu.Unlock()
	if !ok {
		jsonErr(w, http.StatusNotFound, "user not found")
		return
	}
	jsonResp(w, http.StatusOK, map[string]string{"deleted": "ok"})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handleListUsers)
	mux.HandleFunc("GET /users/{id}", handleGetUser)
	mux.HandleFunc("POST /users", handleCreateUser)
	mux.HandleFunc("DELETE /users/{id}", handleDeleteUser)

	slog.Info("flat server on :8081")
	http.ListenAndServe(":8081", loggingMiddleware(mux))
}
