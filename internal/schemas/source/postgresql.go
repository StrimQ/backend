package schemas

// PostgreSQLSourceCreateConfig represents the configuration for a PostgreSQL source.
type PostgreSQLSourceCreateConfig struct {
	Host                string                         `json:"host" validate:"required,hostname"`
	Port                *int                           `json:"port"`
	Username            string                         `json:"username" validate:"required"`
	Password            string                         `json:"password" validate:"required"`
	Database            string                         `json:"database" validate:"required"`
	SnapshotTableSchema *string                        `json:"snapshot_table_schema"`
	SlotName            string                         `json:"slot_name" validate:"required"`
	PublicationName     string                         `json:"publication_name" validate:"required"`
	BinaryHandlingMode  string                         `json:"binary_handling_mode" validate:"omitempty,oneof=bytes base64 base64-url-safe hex"`
	HeartbeatEnabled    bool                           `json:"heartbeat_enabled"`
	HeartbeatInterval   *int                           `json:"heartbeat_interval"`
	HeartbeatSchema     *string                        `json:"heartbeat_schema"`
	HeartbeatTable      *string                        `json:"heartbeat_table"`
	TableHierarchy      map[string]map[string][]string `json:"table_hierarchy" validate:"required"`
}

// SetDefault applies default values to the configuration.
func (c *PostgreSQLSourceCreateConfig) SetDefault() error {
	if c.Port == nil {
		defaultPort := 5432
		c.Port = &defaultPort
	}
	if c.SnapshotTableSchema == nil {
		defaultSchema := "public"
		c.SnapshotTableSchema = &defaultSchema
	}
	if c.BinaryHandlingMode == "" {
		c.BinaryHandlingMode = string(Bytes)
	}

	return nil
}
