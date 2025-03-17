package controllers

import (
	"net/http"

	schemas "github.com/StrimQ/backend/internal/schemas/source"
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
	var sourceCreate schemas.SourceCreate
	if err := sourceCreate.FromIOStream(c.validate, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.sourceService.Create(&sourceCreate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (c *SourceController) Get(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Delete(w http.ResponseWriter, r *http.Request) {
}
