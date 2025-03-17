package schemas

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/go-playground/validator/v10"
)

// SourceConfig defines the interface for source configuration structs.
type SourceConfig interface {
	SetDefault() error
}

// SourceCreate represents the configuration for creating a source.
type SourceCreate struct {
	Name   string       `json:"name" validate:"required"`
	Engine string       `json:"engine" validate:"required,oneof=mysql postgresql"`
	Config SourceConfig `json:"config" validate:"required,dive"`
}

// FromIOStream parses the request body into a SourceCreate struct.
func (s *SourceCreate) FromIOStream(validate *validator.Validate, ioStream io.Reader) error {
	var rawRequest struct {
		Name   string          `json:"name"`
		Engine string          `json:"engine"`
		Config json.RawMessage `json:"config"`
	}
	if err := json.NewDecoder(ioStream).Decode(&rawRequest); err != nil {
		return err
	}
	s.Name = rawRequest.Name
	s.Engine = rawRequest.Engine

	switch s.Engine {
	case "mysql":
		var mysqlConfig MySQLSourceCreateConfig
		if err := json.Unmarshal(rawRequest.Config, &mysqlConfig); err != nil {
			return err
		}
		s.Config = &mysqlConfig
	case "postgresql":
		var pgConfig PostgreSQLSourceCreateConfig
		if err := json.Unmarshal(rawRequest.Config, &pgConfig); err != nil {
			return err
		}
		s.Config = &pgConfig
	default:
		return errors.New("unknown engine type")
	}

	if err := validate.Struct(s); err != nil {
		return err
	}
	if err := s.SetDefault(); err != nil {
		return err
	}

	return nil
}

func (s *SourceCreate) SetDefault() error {
	if err := s.Config.SetDefault(); err != nil {
		return err
	}

	return nil
}
