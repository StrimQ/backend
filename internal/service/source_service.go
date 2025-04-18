package service

import (
	"context"
	"fmt"

	"github.com/StrimQ/backend/internal/client"
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/StrimQ/backend/internal/mapper"
	"github.com/StrimQ/backend/internal/repository"
	"github.com/go-playground/validator/v10"
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
		return nil, fmt.Errorf("validate source request DTO: %w", err)
	}

	// Map the source request DTO to a domain.Source
	source, err := mapper.SourceReqDTOToDomain(ctx, sourceReqDTO)
	if err != nil {
		return nil, fmt.Errorf("map source request DTO to domain: %w", err)
	}

	// Validate the source domain
	if err := source.Validate(s.validate); err != nil {
		return nil, fmt.Errorf("validate source domain: %w", err)
	}

	// Generate collections and topics for the source
	collections, err := source.GenerateCollections()
	if err != nil {
		return nil, fmt.Errorf("generate collections for source: %w", err)
	}

	for i := range collections {
		topic, err := collections[i].GenerateTopic()
		if err != nil {
			return nil, fmt.Errorf("generate topic for collection: %w", err)
		}
		collections[i].Topic = *topic
	}
	source.Collections = collections

	// Generate Kafka Connect configuration
	sourceKCConfig, err := source.GenerateKCConnectorConfig()
	if err != nil {
		return nil, fmt.Errorf("generate Kafka Connect configuration: %w", err)
	}

	// Create the connector in Kafka Connect
	kcClient, ok := ctx.Value(enum.ContextKey_KCClient).(*client.KafkaConnectClient)
	if !ok {
		return nil, fmt.Errorf("failed to get Kafka Connect client from context")
	}
	if err := kcClient.CreateConnnector(ctx, source.GenerateKCConnectorName(), sourceKCConfig); err != nil {
		return nil, fmt.Errorf("create Kafka Connect connector: %w", err)
	}

	// Create the source in the repository
	source, err = s.sourceRepo.Add(ctx, source)
	if err != nil {
		return nil, fmt.Errorf("add source to repository: %w", err)
	}

	// Map the source domain to a source response DTO
	sourceRespDTO, err := mapper.SourceDomainToRespDTO(source)
	if err != nil {
		return nil, fmt.Errorf("map source domain to response DTO: %w", err)
	}

	// Validate the source response DTO
	if err := sourceRespDTO.Validate(s.validate); err != nil {
		return nil, fmt.Errorf("validate source response DTO: %w", err)
	}

	return sourceRespDTO, nil
}
