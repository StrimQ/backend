package mapper

import (
	"encoding/json"
	"errors"

	"github.com/StrimQ/backend/internal/domain"
	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

// SourceReqDTOToDomain maps a SourceDTO to a domain.Source.
func SourceReqDTOToDomain(user *domain.User, sourceID uuid.UUID, sourceReqDTO *dto.SourceReqDTO) (domain.Source, error) {
	// Create source metadata
	metadata := &domain.SourceMetadata{
		TenantID: user.TenantID,
		SourceID: sourceID,
		Name:     sourceReqDTO.Name,
		Engine:   sourceReqDTO.Engine,
	}

	// Create source based on engine
	switch sourceReqDTO.Engine {
	case enum.SourceEngine_Mysql:
		var config domain.MySQLSourceConfig
		if err := json.Unmarshal(sourceReqDTO.Config, &config); err != nil {
			return nil, err
		}
		source := domain.NewMySQLSource(metadata, &config)
		return source, nil
	case enum.SourceEngine_Postgresql:
		var config domain.PostgreSQLSourceConfig
		if err := json.Unmarshal(sourceReqDTO.Config, &config); err != nil {
			return nil, err
		}
		source := domain.NewPostgreSQLSource(metadata, &config)
		return source, nil
	default:
		return nil, errors.New("invalid source engine")
	}
}

func SourceDomainToRespDTO(source domain.Source) (*dto.SourceRespDTO, error) {
	return &dto.SourceRespDTO{
		Name:   source.GetMetadata().Name,
		Engine: source.GetMetadata().Engine,
		// Config: source.GetConfig(),
		// Status: source.GetStatus(),
	}, nil
}
