package controllers

import (
	"net/http"

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

func (c *SourceController) List(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Get(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Delete(w http.ResponseWriter, r *http.Request) {
}
