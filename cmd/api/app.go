package main

import (
	"net/http"
	"os"

	"github.com/StrimQ/backend/internal/controller"
	"github.com/StrimQ/backend/internal/db"
	"github.com/StrimQ/backend/internal/logging"
	"github.com/StrimQ/backend/internal/repository"
	"github.com/StrimQ/backend/internal/service"

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
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	logging.ConfigureLogging(cfg.Debug)

	pgDB, err := db.NewPostgresDB(cfg.PGHost, cfg.PGPort, cfg.PGUsername, cfg.PGPassword, cfg.PGDBName)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to PostgreSQL")
	}

	// Initialize validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	// Dependency injection
	sourceRepo := repository.NewSourceRepository(pgDB)
	sourceService := service.NewSourceService(sourceRepo)
	sourceController := controller.NewSourceController(sourceService, validate)

	router := chi.NewRouter()
	addRoutes(router, sourceController)

	return &App{
		router:   router,
		validate: validate,
	}
}

func addRoutes(router *chi.Mux, sourceController *controller.SourceController) {
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
	log.Info().Msgf("Server starting on :%s", port)
	if err := http.ListenAndServe(":"+port, a.router); err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}
}
