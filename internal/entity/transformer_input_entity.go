package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransformerInputEntity struct {
	TenantID      uuid.UUID
	TransformerID uuid.UUID
	TopicID       uuid.UUID
	Config        []byte
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
