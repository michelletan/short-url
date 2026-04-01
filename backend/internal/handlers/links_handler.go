package handlers

import (
	"encoding/json"
	"net/http"

	"short-url-backend/internal/dtos"
	"short-url-backend/internal/models"
)

type LinkService interface {
	CreateShortLink(userID int, originalURL string) (*models.Link, error)
	GetUserLinks(userID int) ([]models.Link, error)
}

type LinkHandler struct {
	LinkService LinkService
	BaseURL string
}

func NewLinkHandler(linkService LinkService, baseURL string) *LinkHandler {
	return &LinkHandler{LinkService: linkService, BaseURL: baseURL}
}

// POST /api/links
func (h *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("userID")
    userID, ok := ctxUserID.(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	var req dtos.CreateLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	link, err := h.LinkService.CreateShortLink(userID, req.URL)
	if err != nil {
		http.Error(w, "An error occurred", http.StatusInternalServerError)
        return
	}

	res := dtos.CreateLinkResponse{ URL: LinkToShortURLString(h.BaseURL, link) }
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

// GET /api/links
func (h *LinkHandler) List(w http.ResponseWriter, r *http.Request) {
	ctxUserID := r.Context().Value("userID")
    userID, ok := ctxUserID.(int)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	links, err := h.LinkService.GetUserLinks(userID)
	if err != nil {
		http.Error(w, "An error occurred", http.StatusInternalServerError)
        return
	}
	formatted_links := LinksToShortURLStrings(h.BaseURL, links)

	res := dtos.GetLinksResponse{ URLs: formatted_links }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func LinkToShortURLString(baseURL string, link *models.Link) string {
	return baseURL + "/" + link.ShortCode
}

func LinksToShortURLStrings(baseURL string, links []models.Link) []string {
    urls := make([]string, len(links))
    for i, l := range links {
        urls[i] = LinkToShortURLString(baseURL, &l)
    }
    return urls
}
