package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"short-url-server/internal/auth"
	"short-url-server/internal/links"
	"short-url-server/internal/middleware"
	"short-url-server/internal/redirect"
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
