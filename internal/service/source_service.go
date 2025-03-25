package service

import (
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/repository"
)

type SourceService struct {
	sourceRepo *repository.SourceRepository
}

func NewSourceService(sourceRepo *repository.SourceRepository) *SourceService {
	return &SourceService{sourceRepo}
}

/*
TODO:
- Map the sourceCreate to the Source model
- Call the Create method from the sourceRepo
*/
func (s *SourceService) Create(sourceCreate *dto.SourceDTO) error {
	// mapper := NewSourceMapper()
	return nil
}
