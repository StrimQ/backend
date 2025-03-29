package service

import (
	"context"

	"github.com/StrimQ/backend/internal/client"
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
func (s *SourceService) Create(ctx context.Context, sourceReqDTO *dto.SourceReqDTO) (*dto.SourceRespDTO, error) {
	// Validate the source request DTO
	if err := sourceReqDTO.Validate(s.validate); err != nil {
		return nil, err
	}

	// Get the user from the context
	user := ctx.Value(domain.ContextKey_User).(*domain.User)

	// Generate a new source ID
	sourceID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// Map the source request DTO to a domain.Source
	source, err := mapper.SourceReqDTOToDomain(user, sourceID, sourceReqDTO)
	if err != nil {
		return nil, err
	}

	// Validate the source domain
	if err := source.Validate(s.validate); err != nil {
		return nil, err
	}

	// Derive Kafka Connect configuration
	sourceKCConfig, err := source.DeriveKCConnectorConfig()
	if err != nil {
		return nil, err
	}

	// Create the connector in Kafka Connect
	kcClient := client.NewKafkaConnectClient("http://kafka-connect:8083", 10)
	if err := kcClient.CreateConnnector(ctx, source.GetKCConnectorName(), sourceKCConfig); err != nil {
		return nil, err
	}

	// Create the source in the repository
	source, err = s.sourceRepo.Create(ctx, source)
	if err != nil {
		return nil, err
	}

	// Map the source domain to a source response DTO
	sourceRespDTO, err := mapper.SourceDomainToRespDTO(source)
	if err != nil {
		return nil, err
	}

	// Validate the source response DTO
	if err := sourceRespDTO.Validate(s.validate); err != nil {
		return nil, err
	}

	return sourceRespDTO, nil
}
