package domain

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var _ SourceConfig = (*PostgreSQLSourceConfig)(nil)

// PostgreSQLSourceConfig holds PostgreSQL-specific configuration.
type PostgreSQLSourceConfig struct {
	Hostname                 string                         `json:"hostname" validate:"required,hostname"`
	Port                     int                            `json:"port"`
	Username                 string                         `json:"username" validate:"required"`
	Password                 string                         `json:"password" validate:"required"`
	DBName                   string                         `json:"dbName" validate:"required"`
	SSLMode                  enum.SourceSSLMode             `json:"sslMode" validate:"omitempty,oneof=disable require"`
	SlotName                 string                         `json:"slotName" validate:"required"`
	PublicationName          string                         `json:"publicationName" validate:"required"`
	BinaryHandlingMode       enum.SourceBinaryHandlingMode  `json:"binaryHandlingMode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled         bool                           `json:"heartbeatEnabled"`
	HeartbeatIntervalMinutes int                            `json:"heartbeatIntervalMinutes,omitempty"`
	HeartbeatSchema          string                         `json:"heartbeatSchema,omitempty"`
	HeartbeatTable           string                         `json:"heartbeatTable,omitempty"`
	ReadOnly                 bool                           `json:"readOnly"`
	SnapshotSignalSchema     string                         `json:"snapshotSignalSchema,omitempty"`
	SnapshotSignalTable      string                         `json:"snapshotSignalTable,omitempty"`
	CapturedCollections      map[string]map[string][]string `json:"capturedCollections" validate:"required"`
}

// Validate validates the PostgreSQL source configuration and sets default values.
func (c *PostgreSQLSourceConfig) Validate(validate *validator.Validate) error {
	if err := validate.Struct(c); err != nil {
		return err
	}
	if c.Port == 0 {
		c.Port = 5432
	}
	if c.SSLMode == "" {
		c.SSLMode = enum.SourceSSLMode_Require
	}
	if c.BinaryHandlingMode == "" {
		c.BinaryHandlingMode = enum.SourceBinaryHandlingMode_Bytes
	}

	if c.HeartbeatEnabled {
		if c.HeartbeatIntervalMinutes == 0 {
			c.HeartbeatIntervalMinutes = 5
		}
		if !c.ReadOnly {
			if c.HeartbeatSchema == "" && c.HeartbeatTable != "" {
				c.HeartbeatSchema = "public"
			} else if c.HeartbeatSchema != "" && c.HeartbeatTable == "" {
				c.HeartbeatTable = "strimq_heartbeat"
			}
		} else if c.HeartbeatSchema != "" || c.HeartbeatTable != "" {
			return fmt.Errorf("heartbeat schema and table cannot be set for read-only source")
		}
	} else if c.HeartbeatIntervalMinutes != 0 || c.HeartbeatSchema != "" || c.HeartbeatTable != "" {
		return fmt.Errorf("heartbeat interval, schema, and table cannot be set if heartbeat is disabled")
	}

	if !c.ReadOnly {
		if c.SnapshotSignalSchema == "" && c.SnapshotSignalTable != "" {
			c.SnapshotSignalSchema = "public"
		} else if c.SnapshotSignalSchema != "" && c.SnapshotSignalTable == "" {
			c.SnapshotSignalTable = "strimq_snapshot_signal"
		}
	} else if c.SnapshotSignalSchema != "" || c.SnapshotSignalTable != "" {
		return fmt.Errorf("snapshot signal schema and table cannot be set for read-only source")
	}

	return nil
}

func (c *PostgreSQLSourceConfig) AsBytes() ([]byte, error) {
	return json.Marshal(c)
}

func (c *PostgreSQLSourceConfig) GenerateCollections(tenantID uuid.UUID, sourceID uuid.UUID) ([]*SourceCollection, error) {
	// Generate collections based on the captured collections
	// and their respective columns.
	collections := make([]*SourceCollection, 0)
	for groupName, collWithCols := range c.CapturedCollections {
		for collName, columns := range collWithCols {
			collections = append(collections, NewSourceCollection(
				tenantID,
				sourceID,
				c.DBName,
				groupName,
				collName,
				map[string]any{"columns": columns},
			))
		}
	}
	return collections, nil
}

func (c *PostgreSQLSourceConfig) GenerateKCConnectorConfig() (map[string]string, error) {
	kcConfig := map[string]string{
		"connector.class":             "io.debezium.connector.postgresql.PostgresConnector",
		"database.hostname":           c.Hostname,
		"database.port":               strconv.Itoa(c.Port),
		"database.user":               c.Username,
		"database.password":           c.Password,
		"database.dbname":             c.DBName,
		"database.sslmode":            string(c.SSLMode),
		"slot.name":                   c.SlotName,
		"publication.name":            c.PublicationName,
		"binary.handling.mode":        string(c.BinaryHandlingMode),
		"read.only":                   strconv.FormatBool(c.ReadOnly),
		"snapshot.mode":               "no_data",
		"errors.tolerance":            "none",
		"errors.log.enable":           "true",
		"errors.log.include.messages": "false",
	}

	tableIncludeList := make([]string, 0)
	columnIncludeList := make([]string, 0)
	for group, collections := range c.CapturedCollections {
		for collection, columns := range collections {
			tableIncludeList = append(tableIncludeList, fmt.Sprintf("%s.%s", group, collection))
			for _, column := range columns {
				columnIncludeList = append(columnIncludeList, fmt.Sprintf("%s.%s.%s", group, collection, column))
			}
		}
	}

	if c.HeartbeatEnabled {
		heartbeatIntervalMS := (time.Duration(c.HeartbeatIntervalMinutes) * time.Minute).Milliseconds()
		kcConfig["heartbeat.interval.ms"] = strconv.FormatInt(heartbeatIntervalMS, 10)

		if !c.ReadOnly && c.HeartbeatSchema != "" && c.HeartbeatTable != "" {
			query := fmt.Sprintf("INSERT INTO %s.%s (id, last_heartbeat) VALUES (%d, NOW()) ON CONFLICT (id) DO UPDATE SET last_heartbeat = now()", c.HeartbeatSchema, c.HeartbeatTable, 1)
			kcConfig["heartbeat.action.query"] = query
			tableIncludeList = append(tableIncludeList, fmt.Sprintf("%s.%s", c.HeartbeatSchema, c.HeartbeatTable))
		}
	}
	if !c.ReadOnly && c.SnapshotSignalSchema != "" && c.SnapshotSignalTable != "" {
		kcConfig["signal.data.collection"] = fmt.Sprintf("%s.%s", c.SnapshotSignalSchema, c.SnapshotSignalTable)
		tableIncludeList = append(tableIncludeList, fmt.Sprintf("%s.%s", c.SnapshotSignalSchema, c.SnapshotSignalTable))
	}

	kcConfig["table.include.list"] = strings.Join(tableIncludeList, ",")
	kcConfig["column.include.list"] = strings.Join(columnIncludeList, ",")

	return kcConfig, nil
}
