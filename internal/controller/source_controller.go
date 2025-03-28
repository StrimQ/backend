package controller

import (
	"net/http"

	"github.com/StrimQ/backend/internal/dto"
	"github.com/StrimQ/backend/internal/service"
	"github.com/rs/zerolog/log"
)

type SourceController struct {
	sourceService *service.SourceService
}

func NewSourceController(sourceService *service.SourceService) *SourceController {
	return &SourceController{sourceService}
}

func (c *SourceController) List(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Create(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("Creating source")
	var sourceReqDTO dto.SourceReqDTO
	if err := sourceReqDTO.FromIOStream(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sourceRespDTO, err := c.sourceService.Create(r.Context(), &sourceReqDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := sourceRespDTO.ToIOStream(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Info().Msg("Source created")
}

func (c *SourceController) Get(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Update(w http.ResponseWriter, r *http.Request) {
}

func (c *SourceController) Delete(w http.ResponseWriter, r *http.Request) {
}
