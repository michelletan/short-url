package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"short-url-backend/internal/service"
	"short-url-backend/internal/models"
)

type UserService interface {
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

type AuthHandler struct {
	UserService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

// POST /auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    user, err := h.UserService.Register(req.Username, req.Email, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(struct {
        ID       int    `json:"id"`
        Username string `json:"username"`
        Email    string `json:"email"`
    }{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
    })
}

// POST /auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	user, err := h.UserService.Login(req.Email, req.Password)

	if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(struct {
        ID       int    `json:"id"`
        Username string `json:"username"`
        Email    string `json:"email"`
    }{
        ID:       user.ID,
        Username: user.Username,
        Email:    user.Email,
    })

}

// POST /auth/logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout endpoint\n")
}

// GET /auth/me
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Me endpoint (return logged-in user info)\n")
}