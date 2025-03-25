package dto

import (
	"encoding/json"
	"io"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
)

// SourceDTO represents the configuration for creating a source.
type SourceDTO struct {
	Name   string            `json:"name" validate:"required"`
	Engine enum.SourceEngine `json:"engine" validate:"required,oneof=mysql postgresql"`
	Config json.RawMessage   `json:"config" validate:"required"`
}

// FromIOStream parses the request body into a SourceCreate struct.
func (s *SourceDTO) FromIOStream(validate *validator.Validate, ioStream io.Reader) error {
	if err := json.NewDecoder(ioStream).Decode(s); err != nil {
		return err
	}

	if err := validate.Struct(s); err != nil {
		return err
	}

	return nil
}
