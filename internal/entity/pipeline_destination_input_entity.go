package entity

import (
	"time"

	"github.com/google/uuid"
)

type PipelineDestinationInputEntity struct {
	TenantID   uuid.UUID
	PipelineID uuid.UUID
	TopicID    uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
