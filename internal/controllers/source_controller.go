package controllers

import (
	"github.com/StrimQ/backend/internal/services"
	"github.com/go-playground/validator/v10"
)

type SourceController struct {
	sourceService *services.SourceService
	validate      *validator.Validate
}

func NewSourceController(sourceService *services.SourceService, validate *validator.Validate) *SourceController {
	return &SourceController{sourceService, validate}
}

func (c *SourceController) CreateSource() {
	// TODO
}

func (c *SourceController) GetSources() {
	// TODO
}
