package order

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lax/go-relearn/lessons/12-project-layout/examples/domain-driven/internal/store"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", h.list)
	mux.HandleFunc("GET /users/{id}/orders", h.listByUser)
	mux.HandleFunc("POST /orders", h.create)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	orders, err := h.svc.List()
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, orders)
}

func (h *Handler) listByUser(w http.ResponseWriter, r *http.Request) {
	var userID int
	fmt.Sscanf(r.PathValue("id"), "%d", &userID)
	orders, err := h.svc.ListByUser(userID)
	if err != nil {
		writeErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, orders)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int    `json:"user_id"`
		Item   string `json:"item"`
		Qty    int    `json:"qty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeErr(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	o, err := h.svc.Create(input.UserID, input.Item, input.Qty)
	if err != nil {
		writeErr(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, o)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeErr(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, &store.APIError{Status: status, Message: msg})
}
