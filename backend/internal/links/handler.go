package links

import (
	"encoding/json"
	"net/http"

	"short-url-backend/internal/service"
)

type Handler struct {
	URLService *service.URLService
}

func NewHandler(urlService *service.URLService) *Handler {
	return &Handler{URLService: urlService}
}

// POST /api/links
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: parse request body
	// TODO: call service to create short link

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Create link endpoint",
	})
}

// GET /api/links
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: get user from context (set by auth middleware)
	// TODO: fetch user's links

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "List links endpoint",
	})
}
