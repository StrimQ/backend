package mapper

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// SourceDTOToDomain maps a SourceDTO to a domain.Source.
func SourceDTOToDomain(ctx context.Context, validate *validator.Validate, sourceDTO *dto.SourceDTO) (domain.Source, error) {
	// Extract tenant ID from context
	tenantID, ok := ctx.Value("tenant_id").(uuid.UUID)
	if !ok {
		return nil, errors.New("tenant_id not found in context")
	}

	// Generate a new source ID
	sourceID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// Create source metadata
	metadata := &domain.SourceMetadata{
		TenantID: tenantID,
		SourceID: sourceID,
		Name:     sourceDTO.Name,
		Engine:   sourceDTO.Engine,
	}

	// Create and validate the source based on engine
	switch sourceDTO.Engine {
	case enum.SourceEngine_Mysql:
		var config domain.MySQLSourceConfig
		if err := json.Unmarshal(sourceDTO.Config, &config); err != nil {
			return nil, err
		}
		source := domain.NewMySQLSource(metadata, &config)
		if err := source.Validate(validate); err != nil {
			return nil, err
		}
		return source, nil
	case enum.SourceEngine_Postgresql:
		var config domain.PostgreSQLSourceConfig
		if err := json.Unmarshal(sourceDTO.Config, &config); err != nil {
			return nil, err
		}
		source := domain.NewPostgreSQLSource(metadata, &config)
		if err := source.Validate(validate); err != nil {
			return nil, err
		}
		return source, nil
	default:
		return nil, errors.New("invalid source engine")
	}
}
