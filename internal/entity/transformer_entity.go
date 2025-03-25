package entity

import (
	"time"

	"github.com/google/uuid"
)

type TransformerEntity struct {
	TenantID        uuid.UUID
	TransfomerID    uuid.UUID
	Name            string
	Config          string
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
