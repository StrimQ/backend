package repository

import (
	"context"
	"database/sql"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/rs/zerolog/log"
)

type SourceRepository struct {
	db *sql.DB
}

func NewSourceRepository(db *sql.DB) *SourceRepository {
	return &SourceRepository{db}
}

func (r *SourceRepository) Create(ctx context.Context, source *domain.Source) error {
	log.Info().Msgf("Creating source %v", source)
	return nil
}
