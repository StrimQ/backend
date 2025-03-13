package schemas

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// MySQLSourceCreateConfig represents the configuration for a MySQL source.
type MySQLSourceCreateConfig struct {
	Host               string                         `json:"host" validate:"required,hostname"`
	Port               *int                           `json:"port"`
	Database           string                         `json:"database" validate:"required"`
	Username           string                         `json:"username" validate:"required"`
	Password           string                         `json:"password" validate:"required"`
	BinaryHandlingMode string                         `json:"binary_handling_mode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled   bool                           `json:"heartbeat_enabled"`
	HeartbeatInterval  *int                           `json:"heartbeat_interval"`
	HeartbeatSchema    *string                        `json:"heartbeat_schema"`
	HeartbeatTable     *string                        `json:"heartbeat_table"`
	TableHierarchy     map[string]map[string][]string `json:"table_hierarchy" validate:"required"`
}

// Validate validates the MySQLSourceCreateConfig struct, applying defaults and custom logic.
func (c *MySQLSourceCreateConfig) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(c); err != nil {
		return err
	}

	// Apply default values
	if c.Port == nil {
		defaultPort := 3306
		c.Port = &defaultPort
	}
	if c.BinaryHandlingMode == "" {
		c.BinaryHandlingMode = string(Bytes)
	}

	// Validate heartbeat fields
	if c.HeartbeatEnabled {
		if c.HeartbeatInterval == nil {
			return errors.New("heartbeat_interval is required when heartbeat_enabled is true")
		}
		if c.HeartbeatSchema == nil {
			return errors.New("heartbeat_schema is required when heartbeat_enabled is true")
		}
		if c.HeartbeatTable == nil {
			return errors.New("heartbeat_table is required when heartbeat_enabled is true")
		}
	}
	return nil
}
