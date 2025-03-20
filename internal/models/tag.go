package models

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	TenantID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	TagID           uuid.UUID `gorm:"primaryKey;type:uuid"`
	Key             string    `gorm:"type:varchar(255)"`
	Value           string    `gorm:"type:varchar(255)"`
	CreatedByUserID uuid.UUID `gorm:"type:uuid"`
	UpdatedByUserID uuid.UUID `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Associations
	Tenant    Tenant `gorm:"foreignKey:TenantID"`
	CreatedBy User   `gorm:"foreignKey:CreatedByUserID"`
	UpdatedBy User   `gorm:"foreignKey:UpdatedByUserID"`
}
