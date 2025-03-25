package entity

import (
	"time"

	"github.com/google/uuid"
)

type TagEntity struct {
	TenantID        uuid.UUID
	TagID           uuid.UUID
	Key             string
	Value           string
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
