package handlers

import (
	"encoding/json"
	"net/http"

	"short-url-backend/internal/dtos"
	"short-url-backend/internal/models"
)

type LinkService interface {
	CreateShortLink(userID int, originalURL string) (*models.Link, error)
	GetUserLinks(userID int) ([]*models.Link, error)
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

	res := dtos.CreateLinkResponse{ Link: h.LinkToDTO(link) }
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

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

    res := dtos.GetLinksResponse{Links: h.LinksToDTO(links)}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func (h *LinkHandler) LinkToDTO(link *models.Link) dtos.LinkDTO {
    return dtos.LinkDTO{
        ID:         link.ID,
        URL:        link.LongURL,
        Slug:       link.ShortCode,
		ClickCount: link.ClickCount,
        CreatedAt:  link.CreatedAt,
    }
}

func (h *LinkHandler) LinksToDTO(links []*models.Link) []dtos.LinkDTO {
    result := make([]dtos.LinkDTO, len(links))
    for i, l := range links {
        result[i] = h.LinkToDTO(l)
    }
    return result
}