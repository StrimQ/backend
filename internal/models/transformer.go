package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Transformer struct {
	TenantID        uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TransformerID   uuid.UUID       `gorm:"primaryKey;type:uuid"`
	Name            string          `gorm:"type:varchar(255)"`
	Config          json.RawMessage `gorm:"type:jsonb"` // Flink job config, transformation logic
	CreatedByUserID uuid.UUID       `gorm:"type:uuid"`
	UpdatedByUserID uuid.UUID       `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Associations
	Tenant    Tenant              `gorm:"foreignKey:TenantID"`
	Inputs    []TransformerInput  `gorm:"foreignKey:TenantID,TransformerID"`
	Outputs   []TransformerOutput `gorm:"foreignKey:TenantID,TransformerID"`
	CreatedBy User                `gorm:"foreignKey:CreatedByUserID"`
	UpdatedBy User                `gorm:"foreignKey:UpdatedByUserID"`
}

type TransformerInput struct {
	TenantID      uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TransformerID uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TopicID       uuid.UUID       `gorm:"primaryKey;type:uuid"`
	Config        json.RawMessage `gorm:"type:jsonb"` // Input-specific config
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Associations
	Topic Topic `gorm:"foreignKey:TenantID,TopicID"`
}

type TransformerOutput struct {
	TenantID            uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TransformerID       uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TransformerOutputID uuid.UUID       `gorm:"primaryKey;type:uuid"`
	TopicID             uuid.UUID       `gorm:"type:uuid"`  // Output topic produced by transformer
	Config              json.RawMessage `gorm:"type:jsonb"` // Output-specific config
	CreatedAt           time.Time
	UpdatedAt           time.Time

	// Associations
	Topic Topic `gorm:"foreignKey:TenantID,TopicID"`
}
