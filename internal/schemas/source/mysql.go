package schemas

// MySQLSourceCreateConfig represents the configuration for a MySQL source.
type MySQLSourceCreateConfig struct {
	Host               string                         `json:"host" validate:"required,hostname"`
	Port               *int                           `json:"port"`
	Database           string                         `json:"database" validate:"required"`
	Username           string                         `json:"username" validate:"required"`
	Password           string                         `json:"password" validate:"required"`
	BinaryHandlingMode string                         `json:"binary_handling_mode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled   bool                           `json:"heartbeat_enabled"`
	HeartbeatInterval  *int                           `json:"heartbeat_interval" validate:"required_with=HeartbeatEnabled"`
	HeartbeatSchema    *string                        `json:"heartbeat_schema" validate:"required_with=HeartbeatEnabled"`
	HeartbeatTable     *string                        `json:"heartbeat_table" validate:"required_with=HeartbeatEnabled"`
	TableHierarchy     map[string]map[string][]string `json:"table_hierarchy" validate:"required"`
}

// SetDefault applies default values to the configuration.
func (c *MySQLSourceCreateConfig) SetDefault() error {
	if c.Port == nil {
		defaultPort := 3306
		c.Port = &defaultPort
	}
	if c.BinaryHandlingMode == "" {
		c.BinaryHandlingMode = string(Bytes)
	}

	return nil
}
