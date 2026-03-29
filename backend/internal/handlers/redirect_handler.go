package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"short-url-backend/internal/service"
)

type RedirectService interface {
    GetOriginalURL(slug string) (string, error)
	TrackRedirect(urlID int, userIP, userAgent, referrer string) (error)
}

type RedirectHandler struct {
	RedirectService *service.RedirectService
}

func NewRedirectHandler(redirectService *service.RedirectService) *RedirectHandler {
	return &RedirectHandler{RedirectService: redirectService}
}

// GET /{slug}
func (h *RedirectHandler) Handle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	// TODO: look up slug in database
	// TODO: get original URL

	// Temporary placeholder
	originalURL := "https://example.com/" + slug

	http.Redirect(w, r, originalURL, http.StatusFound)
}
