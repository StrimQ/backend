package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Source struct {
	TenantID        uuid.UUID       `gorm:"primaryKey;type:uuid"`
	SourceID        uuid.UUID       `gorm:"primaryKey;type:uuid"`
	Name            string          `gorm:"type:varchar(255)"`
	Engine          SourceEngine    `gorm:"type:varchar(255)"`
	Config          json.RawMessage `gorm:"type:jsonb"` // Connection details, streaming settings
	CreatedByUserID uuid.UUID       `gorm:"type:uuid"`
	UpdatedByUserID uuid.UUID       `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Associations
	Outputs   []SourceOutput `gorm:"foreignKey:TenantID,SourceID"`
	Tenant    Tenant         `gorm:"foreignKey:TenantID"`
	CreatedBy User           `gorm:"foreignKey:CreatedByUserID"`
	UpdatedBy User           `gorm:"foreignKey:UpdatedByUserID"`
}

type SourceOutput struct {
	TenantID       uuid.UUID       `gorm:"primaryKey;type:uuid"`
	SourceID       uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TopicID        uuid.UUID       `gorm:"primaryKey;type:uuid"`
	DatabaseName   string          `gorm:"type:varchar(255)"` // database
	GroupName      string          `gorm:"type:varchar(255)"` // schema/namespace
	CollectionName string          `gorm:"type:varchar(255)"` // /collection/table
	Config         json.RawMessage `gorm:"type:jsonb"`        // Streaming-specific config
	CreatedAt      time.Time
	UpdatedAt      time.Time

	// Associations
	Topic Topic `gorm:"foreignKey:TenantID,TopicID"`
}
