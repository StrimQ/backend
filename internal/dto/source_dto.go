package dto

import (
	"encoding/json"
	"io"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/go-playground/validator/v10"
)

// SourceReqDTO represents the configuration for creating a source.
type SourceReqDTO struct {
	Name   string            `json:"name" validate:"required"`
	Engine enum.SourceEngine `json:"engine" validate:"required,oneof=mysql postgresql"`
	Config json.RawMessage   `json:"config" validate:"required"`
}

// FromIOStream parses the request body into a SourceCreate struct.
func (s *SourceReqDTO) FromIOStream(ioStream io.Reader) error {
	if err := json.NewDecoder(ioStream).Decode(s); err != nil {
		return err
	}

	return nil
}

func (s *SourceReqDTO) Validate(validate *validator.Validate) error {
	return validate.Struct(s)
}

type SourceRespDTO struct {
	Name   string            `json:"name"`
	Engine enum.SourceEngine `json:"engine"`
	Config json.RawMessage   `json:"config"`
	Status json.RawMessage   `json:"status"`
}

// ToIOStream writes the SourceCreate struct into the response body.
func (s *SourceRespDTO) ToIOStream(ioStream io.Writer) error {
	if err := json.NewEncoder(ioStream).Encode(s); err != nil {
		return err
	}

	return nil
}

func (s *SourceRespDTO) Validate(validate *validator.Validate) error {
	return validate.Struct(s)
}
