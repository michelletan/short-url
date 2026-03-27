package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"short-url-backend/internal/auth"
	"short-url-backend/internal/links"
	"short-url-backend/internal/middleware"
	"short-url-backend/internal/redirect"
)

func main() {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", auth.Register)
		r.Post("/login", auth.Login)
		r.Post("/logout", auth.Logout)
		r.Get("/me", auth.Me)
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Route("/links", func(r chi.Router) {
			r.Post("/", links.Create)
			r.Get("/", links.List)
		})
	})

	r.Get("/{slug}", redirect.Handle)

	http.ListenAndServe(":8080", r)
}
