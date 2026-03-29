package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	authMiddleware "short-url-backend/internal/middleware"
	"short-url-backend/internal/service"
	"short-url-backend/internal/store"
	"short-url-backend/internal/db"
	"short-url-backend/internal/handlers"
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
    linkStore := store.NewLinkStore(database)
    redirectStore := store.NewRedirectStore(database)

    // Initialize services
    userService := service.NewUserService(userStore)
    linkService := service.NewLinkService(linkStore)
    redirectService := service.NewRedirectService(redirectStore)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService)
	linksHandler := handlers.NewLinkHandler(linkService)
	redirectHandler := handlers.NewRedirectHandler(redirectService)

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
