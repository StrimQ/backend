package entity

import (
	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

type DestinationEntity struct {
	TenantID        uuid.UUID
	DestinationID   uuid.UUID
	Name            string
	Engine          enum.DestinationEngine
	Config          []byte
	CreatedByUserID uuid.UUID
	UpdatedByUserID uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
