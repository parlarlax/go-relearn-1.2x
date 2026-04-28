package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/db"
)

type Handler struct {
	store *db.Store[*User]
}

func NewHandler(store *db.Store[*User]) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", h.list)
	mux.HandleFunc("GET /users/{id}", h.get)
	mux.HandleFunc("POST /users", h.create)
	mux.HandleFunc("DELETE /users/{id}", h.delete)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.store.List())
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Sscanf(r.PathValue("id"), "%d", &id)
	u, ok := h.store.Get(id)
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, u)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name and email required"})
		return
	}
	u := New(input.Name, input.Email)
	id := h.store.Create(u)
	u.ID = id
	writeJSON(w, http.StatusCreated, u)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Sscanf(r.PathValue("id"), "%d", &id)
	if !h.store.Delete(id) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"deleted": "ok"})
}
