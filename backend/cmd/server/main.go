package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"short-url-backend/internal/auth"
	"short-url-backend/internal/links"
	authMiddleware "short-url-backend/internal/middleware"
	"short-url-backend/internal/redirect"
	"short-url-backend/internal/service"
	"short-url-backend/internal/store"
	"short-url-backend/internal/db"
)

func main() {
	database, err := db.Connect()
    if err != nil {
        log.Fatalf("failed to connect to DB: %v", err)
    }
    defer database.Close()

	log.Println("Handlers initialized, starting server...")

    // Initialize stores
    userStore := store.NewUserStore(database)
    urlStore := store.NewURLStore(database)
    redirectStore := store.NewRedirectStore(database)

    // Initialize services
    userService := service.NewUserService(userStore)
    urlService := service.NewURLService(urlStore)
    redirectService := service.NewRedirectService(redirectStore)

	// Initialize handlers
	authHandler := auth.NewHandler(userService)
	linksHandler := links.NewHandler(urlService)
	redirectHandler := redirect.NewHandler(redirectService)

	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Post("/logout", authHandler.Logout)
		r.Get("/me", authHandler.Me)
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(authMiddleware.Auth)

		r.Route("/links", func(r chi.Router) {
			r.Post("/", linksHandler.Create)
			r.Get("/", linksHandler.List)
		})
	})

	r.Get("/{slug}", redirectHandler.Handle)

	http.ListenAndServe(":8080", r)
}
