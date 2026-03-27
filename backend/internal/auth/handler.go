package auth

import (
	"fmt"
	"net/http"

	"short-url-backend/internal/service"
)

type Handler struct {
	UserService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{UserService: userService}
}

// POST /auth/register
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: parse request body for username, email, password
	// For demo, we'll hardcode a user
	user, err := h.UserService.Register("alice", "alice@example.com", "password123")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Registered user ID: %d\n", user.ID)
}

// POST /auth/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: parse login request
	// Example placeholder: just say login successful
	fmt.Fprintf(w, "Login endpoint (call UserService here)\n")
}

// POST /auth/logout
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout endpoint\n")
}

// GET /auth/me
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me endpoint (return logged-in user info)\n")
}