package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"short-url-backend/internal/config"
	middleware "short-url-backend/internal/middleware"
	"short-url-backend/internal/service"
	"short-url-backend/internal/store"
	"short-url-backend/internal/db"
	"short-url-backend/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

	database, err := db.Connect(cfg.DBURL)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }
    defer database.Close()

	log.Println("Handlers initialized, starting server...")

    // Initialize stores
    userStore := store.NewUserStore(database)
    linkStore := store.NewLinkStore(database)
    redirectStore := store.NewRedirectStore(database)

    // Initialize services
    linkService := service.NewLinkService(linkStore)
    redirectService := service.NewRedirectService(redirectStore)
    jwtService := service.NewJWTService(cfg.JWTSecret, cfg.JWTTTL)
	userService := service.NewUserService(userStore, jwtService)

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
		r.Use(middleware.AuthMiddleware(jwtService))

		r.Route("/links", func(r chi.Router) {
			r.Post("/", linksHandler.Create)
			r.Get("/", linksHandler.List)
		})
	})

	r.Get("/{slug}", redirectHandler.Handle)

	http.ListenAndServe(":8080", r)
}
