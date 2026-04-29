package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lax/go-relearn/examples/third-party/gontainer/layer/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", h.list)
	mux.HandleFunc("GET /users/{id}", h.get)
	mux.HandleFunc("POST /users", h.create)
}

func (h *UserHandler) list(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.List()
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, users)
}

func (h *UserHandler) get(w http.ResponseWriter, r *http.Request) {
	var id int
	if _, err := fmt.Sscanf(r.PathValue("id"), "%d", &id); err != nil || id <= 0 {
		writeErr(w, 400, "invalid id")
		return
	}
	u, err := h.svc.Get(id)
	if err != nil {
		writeErr(w, 404, err.Error())
		return
	}
	writeJSON(w, 200, u)
}

func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeErr(w, 400, "invalid JSON")
		return
	}
	u, err := h.svc.Create(input.Name, input.Email)
	if err != nil {
		writeErr(w, 422, err.Error())
		return
	}
	writeJSON(w, 201, u)
}

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(svc *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: svc}
}

func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", h.list)
	mux.HandleFunc("GET /users/{id}/orders", h.listByUser)
	mux.HandleFunc("POST /orders", h.create)
}

func (h *OrderHandler) list(w http.ResponseWriter, r *http.Request) {
	orders, err := h.svc.List()
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, orders)
}

func (h *OrderHandler) listByUser(w http.ResponseWriter, r *http.Request) {
	var userID int
	fmt.Sscanf(r.PathValue("id"), "%d", &userID)
	orders, err := h.svc.ListByUser(userID)
	if err != nil {
		writeErr(w, 500, err.Error())
		return
	}
	writeJSON(w, 200, orders)
}

func (h *OrderHandler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int    `json:"user_id"`
		Item   string `json:"item"`
		Qty    int    `json:"qty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeErr(w, 400, "invalid JSON")
		return
	}
	o, err := h.svc.Create(input.UserID, input.Item, input.Qty)
	if err != nil {
		writeErr(w, 422, err.Error())
		return
	}
	writeJSON(w, 201, o)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
