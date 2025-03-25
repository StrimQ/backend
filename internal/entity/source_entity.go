package entity

import (
	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

type SourceEntity struct {
	TenantID        uuid.UUID
	SourceID        uuid.UUID
	Name            string
	Engine          enum.SourceEngine
	Config          string
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
