package domain

import (
	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
)

type MySQLSource struct {
	Metadata *SourceMetadata    `validate:"required"`
	Config   *MySQLSourceConfig `validate:"required"`
}

// NewMySQLSource creates a new MySQLSource instance.
func NewMySQLSource(metadata *SourceMetadata, config *MySQLSourceConfig) *MySQLSource {
	return &MySQLSource{
		Metadata: metadata,
		Config:   config,
	}
}

// Validate validates the MySQL source and sets default values.
func (s *MySQLSource) Validate(validate *validator.Validate) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	if err := s.Metadata.Validate(validate); err != nil {
		return err
	}
	if err := s.Config.Validate(validate); err != nil {
		return err
	}

	config := s.Config
	if config.Port == nil {
		defaultPort := 3306
		config.Port = &defaultPort
	}
	if config.BinaryHandlingMode == "" {
		config.BinaryHandlingMode = enum.SourceBinaryHandlingMode_Bytes
	}

	return nil
}

func (s *MySQLSource) GetMetadata() *SourceMetadata {
	return s.Metadata
}

func (s *MySQLSource) GetConfig() SourceConfig {
	return s.Config
}

// DeriveOutputs generates outputs based on the MySQL configuration.
func (s *MySQLSource) DeriveOutputs() ([]SourceOutput, error) {
	var outputs []SourceOutput
	for group, collections := range s.Config.TableHierarchy {
		for collection, columns := range collections {
			outputs = append(outputs, SourceOutput{
				TenantID:       s.Metadata.TenantID,
				SourceID:       s.Metadata.SourceID,
				DatabaseName:   s.Config.Database,
				GroupName:      group,
				CollectionName: collection,
				Config:         map[string]any{"columns": columns},
			})
		}
	}
	return outputs, nil
}

// MySQLSourceConfig holds MySQL-specific configuration.
type MySQLSourceConfig struct {
	Host               string                         `json:"host" validate:"required,hostname"`
	Port               *int                           `json:"port"`
	Database           string                         `json:"database" validate:"required"`
	Username           string                         `json:"username" validate:"required"`
	Password           string                         `json:"password" validate:"required"`
	BinaryHandlingMode enum.SourceBinaryHandlingMode  `json:"binary_handling_mode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled   bool                           `json:"heartbeat_enabled"`
	HeartbeatInterval  *int                           `json:"heartbeat_interval" validate:"required_with=HeartbeatEnabled"`
	HeartbeatSchema    *string                        `json:"heartbeat_schema" validate:"required_with=HeartbeatEnabled"`
	HeartbeatTable     *string                        `json:"heartbeat_table" validate:"required_with=HeartbeatEnabled"`
	TableHierarchy     map[string]map[string][]string `json:"table_hierarchy" validate:"required"`
}

// Validate validates the MySQL source configuration.
func (c *MySQLSourceConfig) Validate(validate *validator.Validate) error {
	return validate.Struct(c)
}
