package mapper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

// SourceReqDTOToDomain maps a SourceReqDTO to a domain.Source.
func SourceReqDTOToDomain(ctx context.Context, sourceReqDTO *dto.SourceReqDTO) (*domain.Source, error) {
	// Extract user from context
	user, ok := ctx.Value(enum.ContextKey_User).(*domain.User)
	if !ok || user == nil {
		return nil, fmt.Errorf("user not found in context")
	}

	// Generate a new UUID for the source
	sourceID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("generate source ID: %w", err)
	}

	// Determine and unmarshal the config based on the engine
	var config domain.SourceConfig
	switch sourceReqDTO.Engine {
	case enum.SourceEngine_Mysql:
		var mysqlConfig domain.MySQLSourceConfig
		if err := json.Unmarshal(sourceReqDTO.Config, &mysqlConfig); err != nil {
			return nil, fmt.Errorf("unmarshal MySQL config: %w", err)
		}
		config = &mysqlConfig
	case enum.SourceEngine_Postgresql:
		var pgConfig domain.PostgreSQLSourceConfig
		if err := json.Unmarshal(sourceReqDTO.Config, &pgConfig); err != nil {
			return nil, fmt.Errorf("unmarshal PostgreSQL config: %w", err)
		}
		config = &pgConfig
	default:
		return nil, fmt.Errorf("invalid source engine: %s", sourceReqDTO.Engine)
	}

	return domain.NewSource(
		user.TenantID,
		sourceID,
		sourceReqDTO.Name,
		sourceReqDTO.Engine,
		config,
		user.UserID,
		user.UserID,
	), nil
}

// SourceDomainToRespDTO maps a domain.Source to a SourceRespDTO.
func SourceDomainToRespDTO(source *domain.Source) (*dto.SourceRespDTO, error) {
	configBytes, err := source.Config.AsBytes()
	if err != nil {
		return nil, fmt.Errorf("serialize source config: %w", err)
	}

	return &dto.SourceRespDTO{
		ID:     source.SourceID,
		Name:   source.Name,
		Engine: source.Engine,
		Config: configBytes,
	}, nil
}
