package service

import (
	"context"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/mapper"
	"github.com/StrimQ/backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SourceService struct {
	validate   *validator.Validate
	sourceRepo *repository.SourceRepository
}

func NewSourceService(validate *validator.Validate, sourceRepo *repository.SourceRepository) *SourceService {
	return &SourceService{validate, sourceRepo}
}

/*
TODO:
- Map the sourceCreate to the Source model
- Call the Create method from the sourceRepo
*/
func (s *SourceService) Create(ctx context.Context, sourceDTO *dto.SourceDTO) error {
	// Get the user from the context
	user := ctx.Value(domain.ContextKey_User).(*domain.User)

	// Generate a new source ID
	sourceID, err := uuid.NewV7()
	if err != nil {
		return err
	}

	source, err := mapper.SourceDTOToDomain(user, sourceID, sourceDTO)
	if err != nil {
		return err
	}

	if err := source.Validate(s.validate); err != nil {
		return err
	}

	if err := s.sourceRepo.Create(ctx, &source); err != nil {
		return err
	}

	return nil
}
