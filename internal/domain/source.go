package domain

import (
	"strings"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Source defines the common behavior for all source types.
type Source interface {
	Validate(validate *validator.Validate) error
	GetMetadata() *SourceMetadata
	GetConfig() SourceConfig
	DeriveOutputs() ([]SourceOutput, error)
	DeriveKCConfig() (map[string]string, error)
}

type SourceConfig interface{}

type SourceMetadata struct {
	TenantID uuid.UUID         `validate:"required"`
	SourceID uuid.UUID         `validate:"required"`
	Name     string            `validate:"required"`
	Engine   enum.SourceEngine `validate:"required,oneof=mysql postgresql"`
}

func (m *SourceMetadata) Validate(validate *validator.Validate) error {
	return validate.Struct(m)
}

type SourceOutput struct {
	TenantID       uuid.UUID `validate:"required"`
	SourceID       uuid.UUID `validate:"required"`
	DatabaseName   string
	GroupName      string
	CollectionName string `validate:"required"`
	Config         map[string]any
}

func (o *SourceOutput) DeriveTopic() (*Topic, error) {
	topicID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// Derive the topic name from the output
	nameParts := []string{
		o.TenantID.String(),
		o.SourceID.String(),
	}
	if o.DatabaseName != "" {
		nameParts = append(nameParts, o.DatabaseName)
	}
	if o.GroupName != "" {
		nameParts = append(nameParts, o.GroupName)
	}
	nameParts = append(nameParts, o.CollectionName)
	topicName := strings.Join(nameParts, ".")

	return &Topic{
		TenantID: o.TenantID,
		TopicID:  topicID,
		Name:     topicName,
	}, nil
}
