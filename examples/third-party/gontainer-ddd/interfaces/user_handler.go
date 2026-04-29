package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lax/go-relearn/examples/third-party/gontainer-ddd/application"
)

type UserHandler struct {
	svc *application.UserService
}

func NewUserHandler(svc *application.UserService) *UserHandler {
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
