package models

import (
	"time"

	"github.com/google/uuid"
)

type Topic struct {
	TenantID     uuid.UUID         `gorm:"primaryKey;type:uuid"`
	TopicID      uuid.UUID         `gorm:"primaryKey;type:uuid"`
	Name         string            `gorm:"type:varchar(255)"` // Kafka topic name, e.g., "tenant1.source1.collectionA"
	ProducerType TopicProducerType `gorm:"type:varchar(255)"`
	ProducerID   uuid.UUID         `gorm:"type:uuid"` // References sources.source_id or transformers.transformer_id based on producer_type
	CreatedAt    time.Time
	UpdatedAt    time.Time

	// Associations
	Tenant Tenant `gorm:"foreignKey:TenantID"`
}
