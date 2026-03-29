package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"

	"short-url-backend/internal/models"
	"short-url-backend/internal/dtos"
)

type UserService interface {
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (dtos.LoginResponse, error)
	GetByID(id int) (*models.User, error)
}

type AuthHandler struct {
	UserService UserService
}

func NewAuthHandler(userService UserService) *AuthHandler {
	return &AuthHandler{UserService: userService}
}

// POST /auth/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    user, err := h.UserService.Register(req.Username, req.Email, req.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	res := dtos.RegisterResponse{
		Message: fmt.Sprintf("User %s registered successfully", user.Username),
	}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

// POST /auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.LoginRequest

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	tokens, err := h.UserService.Login(req.Email, req.Password)

	if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tokens)
}

// POST /auth/logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GET /auth/me
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(int)
    user, err := h.UserService.GetByID(userID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

	var res = dtos.MeResponse{
		Email:    user.Email,
		Username: user.Username,
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}