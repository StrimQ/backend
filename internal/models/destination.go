package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Destination struct {
	TenantID        uuid.UUID         `gorm:"primaryKey;type:uuid"`
	DestinationID   uuid.UUID         `gorm:"primaryKey;type:uuid"`
	Name            string            `gorm:"type:varchar(255)"`
	Engine          DestinationEngine `gorm:"type:varchar(255)"` // Enum for destination type
	Config          json.RawMessage   `gorm:"type:jsonb"`        // Connection details, write settings
	CreatedByUserID uuid.UUID         `gorm:"type:uuid"`
	UpdatedByUserID uuid.UUID         `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Associations
	Tenant    Tenant `gorm:"foreignKey:TenantID"`
	CreatedBy User   `gorm:"foreignKey:CreatedByUserID"`
	UpdatedBy User   `gorm:"foreignKey:UpdatedByUserID"`
}
