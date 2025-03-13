package schemas

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
)

// SourceConfig defines the interface for source configuration structs.
type SourceConfig interface {
	Validate() error
}

// SourceCreate represents the configuration for creating a source.
type SourceCreate struct {
	Name      string       `json:"name" validate:"required"`
	Connector string       `json:"connector" validate:"required,oneof=mysql postgresql"`
	Config    SourceConfig `json:"config" validate:"required"`
}

// UnmarshalJSON customizes JSON unmarshaling to handle the discriminated union based on "connector".
func (s *SourceCreate) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name      string          `json:"name"`
		Connector string          `json:"connector"`
		Config    json.RawMessage `json:"config"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.Name = temp.Name
	s.Connector = temp.Connector

	switch s.Connector {
	case "mysql":
		var mysqlConfig MySQLSourceCreateConfig
		if err := json.Unmarshal(temp.Config, &mysqlConfig); err != nil {
			return err
		}
		s.Config = &mysqlConfig
	case "postgresql":
		var pgConfig PostgreSQLSourceCreateConfig
		if err := json.Unmarshal(temp.Config, &pgConfig); err != nil {
			return err
		}
		s.Config = &pgConfig
	default:
		return errors.New("unknown connector type")
	}
	return nil
}

// Validate validates the SourceCreate struct.
func (s *SourceCreate) Validate() error {
	validate := validator.New()
	if err := validate.Struct(s); err != nil {
		return err
	}
	if s.Config == nil { // Shouldn't happen if UnmarshalJSON succeeds, but added for safety
		return errors.New("config is required")
	}
	return s.Config.Validate()
}
