package domain

import (
	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
)

// PostgreSQLSource represents a PostgreSQL data source.
type PostgreSQLSource struct {
	Metadata *SourceMetadata         `validate:"required"`
	Config   *PostgreSQLSourceConfig `validate:"required"`
}

// NewPostgreSQLSource creates a new PostgreSQLSource instance.
func NewPostgreSQLSource(metadata *SourceMetadata, config *PostgreSQLSourceConfig) *PostgreSQLSource {
	return &PostgreSQLSource{
		Metadata: metadata,
		Config:   config,
	}
}

// Validate validates the PostgreSQL source and sets default values.
func (s *PostgreSQLSource) Validate(validate *validator.Validate) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	if err := s.Metadata.Validate(validate); err != nil {
		return err
	}
	if err := s.Config.Validate(validate); err != nil {
		return err
	}

	if s.Config.Port == nil {
		defaultPort := 5432
		s.Config.Port = &defaultPort
	}
	if s.Config.SnapshotTableSchema == nil {
		defaultSchema := "public"
		s.Config.SnapshotTableSchema = &defaultSchema
	}
	if s.Config.BinaryHandlingMode == "" {
		s.Config.BinaryHandlingMode = enum.SourceBinaryHandlingMode_Bytes
	}
	return nil
}

// GenerateOutputs generates outputs based on the PostgreSQL configuration.
func (s *PostgreSQLSource) GenerateOutputs() []SourceOutput {
	var outputs []SourceOutput
	for group, collections := range s.Config.TableHierarchy {
		for collection, columns := range collections {
			outputs = append(outputs, SourceOutput{
				DatabaseName:   s.Config.Database,
				GroupName:      group,
				CollectionName: collection,
				Config:         map[string]any{"columns": columns},
			})
		}
	}
	return outputs
}

// PostgreSQLSourceConfig holds PostgreSQL-specific configuration.
type PostgreSQLSourceConfig struct {
	Host                string                         `json:"host" validate:"required,hostname"`
	Port                *int                           `json:"port"`
	Username            string                         `json:"username" validate:"required"`
	Password            string                         `json:"password" validate:"required"`
	Database            string                         `json:"database" validate:"required"`
	SnapshotTableSchema *string                        `json:"snapshot_table_schema"`
	SlotName            string                         `json:"slot_name" validate:"required"`
	PublicationName     string                         `json:"publication_name" validate:"required"`
	BinaryHandlingMode  enum.SourceBinaryHandlingMode  `json:"binary_handling_mode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled    bool                           `json:"heartbeat_enabled"`
	HeartbeatInterval   *int                           `json:"heartbeat_interval"`
	HeartbeatSchema     *string                        `json:"heartbeat_schema"`
	HeartbeatTable      *string                        `json:"heartbeat_table"`
	TableHierarchy      map[string]map[string][]string `json:"table_hierarchy" validate:"required"`
}

// Validate validates the PostgreSQL source configuration.
func (c *PostgreSQLSourceConfig) Validate(validate *validator.Validate) error {
	return validate.Struct(c)
}
