package mapper

import (
	"encoding/json"
	"errors"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

// SourceDTOToDomain maps a SourceDTO to a domain.Source.
func SourceDTOToDomain(user *domain.User, sourceID uuid.UUID, sourceDTO *dto.SourceDTO) (domain.Source, error) {
	// Create source metadata
	metadata := &domain.SourceMetadata{
		TenantID: user.TenantID,
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
		return source, nil
	case enum.SourceEngine_Postgresql:
		var config domain.PostgreSQLSourceConfig
		if err := json.Unmarshal(sourceDTO.Config, &config); err != nil {
			return nil, err
		}
		source := domain.NewPostgreSQLSource(metadata, &config)
		return source, nil
	default:
		return nil, errors.New("invalid source engine")
	}
}
