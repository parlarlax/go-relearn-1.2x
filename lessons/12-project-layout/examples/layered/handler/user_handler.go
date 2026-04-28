package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/layered/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", h.List)
	mux.HandleFunc("GET /users/{id}", h.Get)
	mux.HandleFunc("POST /users", h.Create)
	mux.HandleFunc("DELETE /users/{id}", h.Delete)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.svc.List())
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Sscanf(r.PathValue("id"), "%d", &id)
	u, err := h.svc.GetByID(id)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, u)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid input"})
		return
	}
	u, err := h.svc.Create(input.Name, input.Email)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, u)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var id int
	fmt.Sscanf(r.PathValue("id"), "%d", &id)
	if err := h.svc.Delete(id); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"deleted": "ok"})
}
