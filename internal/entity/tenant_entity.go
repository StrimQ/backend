package entity

import (
	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

type TenantEntity struct {
	TenantID  uuid.UUID
	Name      string
	Domain    string
	Tier      enum.TenantTier
	InfraID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
