package main

import (
	"net/http"
	"os"

	"github.com/StrimQ/backend/internal/controllers"
	"github.com/StrimQ/backend/internal/db"
	"github.com/StrimQ/backend/internal/logging"
	"github.com/StrimQ/backend/internal/repositories"
	"github.com/StrimQ/backend/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// App holds application components for dependency injection
type App struct {
	router *chi.Mux

	validate *validator.Validate
}

// NewApp initializes the application
func NewApp() *App {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config: %v")
	}

	logging.ConfigureLogging(cfg.Debug)

	pgDB, err := db.NewPostgresDB(cfg.PGHost)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to PostgreSQL")
	}

	// Initialize validator
	validate := validator.New()

	// Dependency injection
	sourceRepo := repositories.NewSourceRepository(pgDB)
	sourceService := services.NewSourceService(sourceRepo)
	sourceController := controllers.NewSourceController(sourceService, validate)

	router := chi.NewRouter()
	addRoutes(router, sourceController)

	return &App{
		router:   router,
		validate: validate,
	}
}

func addRoutes(router *chi.Mux, sourceController *controllers.SourceController) {
	router.Get("/sources", sourceController.List)
	router.Post("/sources", sourceController.Create)
	router.Get("/sources/{id}", sourceController.Get)
	router.Put("/sources/{id}", sourceController.Update)
	router.Delete("/sources/{id}", sourceController.Delete)
}

// Run starts the HTTP server
func (a *App) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, a.router); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
