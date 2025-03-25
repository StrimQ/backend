package domain

import (
	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Source defines the common behavior for all source types.
type Source interface {
	Validate(validate *validator.Validate) error
	GenerateOutputs() []SourceOutput
}

type SourceConfig interface {
}

type SourceMetadata struct {
	TenantID uuid.UUID
	SourceID uuid.UUID
	Name     string            `validate:"required"`
	Engine   enum.SourceEngine `validate:"required,oneof=mysql postgresql"`
}

type SourceOutput struct {
	DatabaseName   string
	GroupName      string
	CollectionName string
	Config         map[string]any
}
