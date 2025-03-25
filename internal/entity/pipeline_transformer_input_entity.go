package entity

import (
	"time"

	"github.com/google/uuid"
)

type PipelineTransformerInputEntity struct {
	TenantID      uuid.UUID
	PipelineID    uuid.UUID
	TransformerID uuid.UUID
	TopicID       uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
