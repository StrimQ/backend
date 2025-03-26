package entity

import (
	"time"

	"github.com/google/uuid"
)

type PipelineEntity struct {
	TenantID        uuid.UUID
	PipelineID      uuid.UUID
	Name            string
	SourceID        *uuid.UUID
	DestinationID   uuid.UUID
	Config          []byte
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
