package controller

import (
	"net/http"

	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/service"
	"github.com/go-playground/validator/v10"
)

type SourceController struct {
	sourceService *service.SourceService
	validate      *validator.Validate
}

func NewSourceController(sourceService *service.SourceService, validate *validator.Validate) *SourceController {
	return &SourceController{sourceService, validate}
}

func (c *SourceController) List(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Create(w http.ResponseWriter, r *http.Request) {
	var sourceDTO dto.SourceDTO
	if err := sourceDTO.FromIOStream(c.validate, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.sourceService.Create(&sourceDTO); err != nil {
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
