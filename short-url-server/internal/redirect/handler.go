package redirect

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GET /{slug}
func Handle(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	// TODO: look up slug in database
	// TODO: get original URL

	// Temporary placeholder
	originalURL := "https://example.com/" + slug

	http.Redirect(w, r, originalURL, http.StatusFound)
}
