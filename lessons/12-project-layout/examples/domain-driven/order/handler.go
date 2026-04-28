package order

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
}

func New(userID int, item string) *Order {
	return &Order{UserID: userID, Item: item}
}

type Handler struct {
	orders []*Order
	nextID int
}

func NewHandler() *Handler {
	return &Handler{nextID: 1}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /orders", h.list)
	mux.HandleFunc("POST /orders", h.create)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, h.orders)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int    `json:"user_id"`
		Item   string `json:"item"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Item == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "item required"})
		return
	}
	o := New(input.UserID, input.Item)
	o.ID = h.nextID
	h.nextID++
	h.orders = append(h.orders, o)
	writeJSON(w, http.StatusCreated, o)
}
