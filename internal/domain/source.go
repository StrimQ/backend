package domain

import (
	"strings"

	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Source defines the common behavior for all source types.
type Source struct {
	TenantID        uuid.UUID
	SourceID        uuid.UUID
	Name            string            `validate:"required"`
	Engine          enum.SourceEngine `validate:"required,oneof=mysql postgresql"`
	Config          SourceConfig
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Associations
	Collections []*SourceCollection
}

func NewSource(
	tenantID uuid.UUID,
	sourceID uuid.UUID,
	name string,
	engine enum.SourceEngine,
	config SourceConfig,
	createdByUserID uuid.UUID,
	updatedByUserID uuid.UUID,
) *Source {
	return &Source{
		TenantID:        tenantID,
		SourceID:        sourceID,
		Name:            name,
		Engine:          engine,
		Config:          config,
		CreatedByUserID: createdByUserID,
		UpdatedByUserID: updatedByUserID,
	}
}

func (s *Source) Validate(validate *validator.Validate) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	if err := s.Config.Validate(validate); err != nil {
		return err
	}
	return nil
}

func (s *Source) GenerateCollections() ([]*SourceCollection, error) {
	return s.Config.GenerateCollections(s.TenantID, s.SourceID)
}

func (s *Source) GenerateKCConnectorName() string {
	return strings.Join([]string{
		s.TenantID.String(),
		s.SourceID.String(),
	}, ".",
	)
}

func (s *Source) GenerateKCConnectorConfig() (map[string]string, error) {
	return s.Config.GenerateKCConnectorConfig(s.GenerateKCConnectorName())
}

type SourceConfig interface {
	Validate(validate *validator.Validate) error
	AsBytes() ([]byte, error)
	GenerateCollections(tenantID uuid.UUID, sourceID uuid.UUID) ([]*SourceCollection, error)
	GenerateKCConnectorConfig(kcConnectorName string) (map[string]string, error)
}

type SourceCollection struct {
	TenantID       uuid.UUID
	SourceID       uuid.UUID
	TopicID        uuid.UUID
	DatabaseName   string
	GroupName      string
	CollectionName string         `validate:"required"`
	Config         map[string]any `validate:"required"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// Associations
	Topic *Topic
}

func NewSourceCollection(
	tenantID uuid.UUID,
	sourceID uuid.UUID,
	databaseName string,
	groupName string,
	collectionName string,
	config map[string]any,
) *SourceCollection {
	return &SourceCollection{
		TenantID:       tenantID,
		SourceID:       sourceID,
		DatabaseName:   databaseName,
		GroupName:      groupName,
		CollectionName: collectionName,
		Config:         config,
	}
}

func (o *SourceCollection) GenerateTopic() (*Topic, error) {
	topicID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	// Generate the topic name from the collection
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
