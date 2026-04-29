package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lax/go-relearn/examples/gontainer-ddd/application"
)

type OrderHandler struct {
	svc *application.OrderService
}

func NewOrderHandler(svc *application.OrderService) *OrderHandler {
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
