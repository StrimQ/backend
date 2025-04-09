package domain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Ensure MySQLSourceConfig implements SourceConfig
var _ SourceConfig = (*MySQLSourceConfig)(nil)

// MySQLSourceConfig holds MySQL-specific configuration.
type MySQLSourceConfig struct {
	ServerName               string                         `json:"serverName" validate:"required"`
	Host                     string                         `json:"host" validate:"required,hostname"`
	Port                     int                            `json:"port"`
	Database                 string                         `json:"database" validate:"required"`
	Username                 string                         `json:"username" validate:"required"`
	Password                 string                         `json:"password" validate:"required"`
	BinaryHandlingMode       enum.SourceBinaryHandlingMode  `json:"binaryHandlingMode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled         bool                           `json:"heartbeatEnabled"`
	HeartbeatIntervalMinutes int                            `json:"heartbeatIntervalMinutes,omitempty"`
	SignalTable              string                         `json:"signalTable,omitempty"`
	CapturedCollections      map[string]map[string][]string `json:"capturedCollections" validate:"required"`
}

// Validate validates the MySQL source configuration and sets default values.
func (c *MySQLSourceConfig) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return err
	}
	if c.Port == 0 {
		c.Port = 3306 // Default MySQL port
	}
	if c.BinaryHandlingMode == "" {
		c.BinaryHandlingMode = enum.SourceBinaryHandlingMode_Bytes
	}
	if c.HeartbeatEnabled && c.HeartbeatIntervalMinutes == 0 {
		c.HeartbeatIntervalMinutes = 5 // Default to 5 minutes
	}
	return nil
}

// AsBytes serializes the configuration to JSON bytes.
func (c *MySQLSourceConfig) AsBytes() ([]byte, error) {
	return json.Marshal(c)
}

// GenerateCollections generates SourceCollection instances based on captured collections.
func (c *MySQLSourceConfig) GenerateCollections(tenantID uuid.UUID, sourceID uuid.UUID) ([]*SourceCollection, error) {
	collections := make([]*SourceCollection, 0)
	for db, tables := range c.CapturedCollections {
		for table, columns := range tables {
			collections = append(collections, NewSourceCollection(
				tenantID,
				sourceID,
				db, // Database name from CapturedCollections
				"", // GroupName not used in MySQL context
				table,
				map[string]any{"columns": columns},
			))
		}
	}
	return collections, nil
}

// GenerateKCConnectorConfig generates Kafka Connect configuration for the MySQL connector.
func (c *MySQLSourceConfig) GenerateKCConnectorConfig() (map[string]string, error) {
	kcConfig := map[string]string{
		"connector.class":                          "io.debezium.connector.mysql.MySqlConnector",
		"database.hostname":                        c.Host,
		"database.port":                            strconv.Itoa(c.Port),
		"database.user":                            c.Username,
		"database.password":                        c.Password,
		"database.server.id":                       "1", // Consider making configurable in production
		"database.server.name":                     c.ServerName,
		"database.history.kafka.bootstrap.servers": "kafka:9092",
		"database.history.kafka.topic":             "schema-changes." + c.ServerName,
		"binary.handling.mode":                     string(c.BinaryHandlingMode),
	}

	// Configure heartbeat if enabled
	if c.HeartbeatEnabled {
		heartbeatIntervalMS := (time.Duration(c.HeartbeatIntervalMinutes) * time.Minute).Milliseconds()
		kcConfig["heartbeat.interval.ms"] = strconv.FormatInt(heartbeatIntervalMS, 10)
	}

	// Configure signal table if specified
	if c.SignalTable != "" {
		kcConfig["signal.data.collection"] = fmt.Sprintf("%s.%s", c.Database, c.SignalTable)
	}

	// Build table.include.list from CapturedCollections
	tableIncludeList := make([]string, 0)
	for db, tables := range c.CapturedCollections {
		for table := range tables {
			tableIncludeList = append(tableIncludeList, fmt.Sprintf("%s.%s", db, table))
		}
	}
	if len(tableIncludeList) > 0 {
		kcConfig["table.include.list"] = strings.Join(tableIncludeList, ",")
	}

	return kcConfig, nil
}
