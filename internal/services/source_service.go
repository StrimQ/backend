package services

import "github.com/StrimQ/backend/internal/repositories"

type SourceService struct {
	sourceRepo *repositories.SourceRepository
}

func NewSourceService(sourceRepo *repositories.SourceRepository) *SourceService {
	return &SourceService{sourceRepo}
}
