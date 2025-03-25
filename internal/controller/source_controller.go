package controller

import (
	"net/http"

	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type SourceController struct {
	validate      *validator.Validate
	sourceService *service.SourceService
}

func NewSourceController(validate *validator.Validate, sourceService *service.SourceService) *SourceController {
	return &SourceController{validate, sourceService}
}

func (c *SourceController) List(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Create(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Creating source")
	var sourceDTO dto.SourceDTO
	if err := sourceDTO.FromIOStream(c.validate, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.sourceService.Create(r.Context(), &sourceDTO); err != nil {
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
