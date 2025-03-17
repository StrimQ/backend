package services

import (
	"github.com/StrimQ/backend/internal/repositories"
	schemas "github.com/StrimQ/backend/internal/schemas/source"
)

type SourceService struct {
	sourceRepo *repositories.SourceRepository
}

func NewSourceService(sourceRepo *repositories.SourceRepository) *SourceService {
	return &SourceService{sourceRepo}
}

/*
TODO: From sourceCreate:
- Extract and transform into PostgreSQL source model
- Extract
*/
func (s *SourceService) Create(sourceCreate *schemas.SourceCreate) error {
	return nil
}
