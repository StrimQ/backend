package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransformerOutputEntity struct {
	TenantID      uuid.UUID
	TransformerID uuid.UUID
	TopicID       uuid.UUID
	Config        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
