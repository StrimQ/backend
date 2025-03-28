package domain

import (
	"strconv"
	"strings"

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
	return nil
}

func (s *MySQLSource) GetMetadata() *SourceMetadata {
	return s.Metadata
}

func (s *MySQLSource) GetConfig() SourceConfig {
	return s.Config
}

func (s *MySQLSource) GetKCConnectorName() string {
	return strings.Join([]string{
		s.Metadata.TenantID.String(),
		s.Metadata.SourceID.String(),
	}, ".",
	)
}

// DeriveOutputs generates outputs based on the MySQL configuration.
func (s *MySQLSource) DeriveOutputs() ([]SourceOutput, error) {
	var outputs []SourceOutput
	for group, collections := range s.Config.CapturedCollections {
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

func (s *MySQLSource) DeriveKCConnectorConfig() (map[string]string, error) {
	return s.Config.DeriveKCConnectorConfig()
}

// MySQLSourceConfig holds MySQL-specific configuration.
type MySQLSourceConfig struct {
	Host                string                         `json:"host" validate:"required,hostname"`
	Port                *int                           `json:"port"`
	Database            string                         `json:"database" validate:"required"`
	Username            string                         `json:"username" validate:"required"`
	Password            string                         `json:"password" validate:"required"`
	BinaryHandlingMode  enum.SourceBinaryHandlingMode  `json:"binaryHandlingMode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled    bool                           `json:"heartbeatEnabled"`
	HeartbeatInterval   *int                           `json:"heartbeatInterval" validate:"required_with=HeartbeatEnabled"`
	HeartbeatSchema     *string                        `json:"heartbeatSchema" validate:"required_with=HeartbeatEnabled"`
	HeartbeatTable      *string                        `json:"heartbeatTable" validate:"required_with=HeartbeatEnabled"`
	CapturedCollections map[string]map[string][]string `json:"capturedCollections" validate:"required"`
}

// Validate validates the MySQL source configuration and sets default values.
func (c *MySQLSourceConfig) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return err
	}
	if c.Port == nil {
		defaultPort := 3306
		c.Port = &defaultPort
	}
	if c.BinaryHandlingMode == "" {
		c.BinaryHandlingMode = enum.SourceBinaryHandlingMode_Bytes
	}
	return nil
}

// DeriveKCConnectorConfig generates Kafka Connect configuration based on the MySQL configuration.
func (c *MySQLSourceConfig) DeriveKCConnectorConfig() (map[string]string, error) {
	return map[string]string{
		"connector.class":                          "io.debezium.connector.mysql.MySqlConnector",
		"database.hostname":                        c.Host,
		"database.port":                            strconv.Itoa(*c.Port),
		"database.user":                            c.Username,
		"database.password":                        c.Password,
		"database.server.id":                       "1",
		"database.server.name":                     c.Database,
		"database.whitelist":                       c.Database,
		"database.history.kafka.bootstrap.servers": "kafka:9092",
		"database.history.kafka.topic":             "schema-changes." + c.Database,
	}, nil
}
