package entity

import (
	"time"

	"github.com/google/uuid"
)

type PipelineTransformerEntity struct {
	TenantID      uuid.UUID
	PipelineID    uuid.UUID
	TransformerID uuid.UUID
	Stage         int32
	Config        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
