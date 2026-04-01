package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"short-url-backend/internal/service"
	"short-url-backend/internal/models"
)

type RedirectService interface {
    GetLinkByShortCode(slug string) (*models.Link, error)
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

	link, err := h.RedirectService.GetLinkByShortCode(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	h.RedirectService.TrackRedirect(link.ID, r.RemoteAddr, r.UserAgent(), r.Referer())

	http.Redirect(w, r, link.LongURL, http.StatusFound)
}
