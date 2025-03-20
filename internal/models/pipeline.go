package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Pipeline struct {
	TenantID        uuid.UUID       `gorm:"primaryKey;type:uuid"`
	PipelineID      uuid.UUID       `gorm:"primaryKey;type:uuid"`
	Name            string          `gorm:"type:varchar(255)"`
	SourceID        *uuid.UUID      `gorm:"type:uuid"` // Nullable field
	DestinationID   uuid.UUID       `gorm:"type:uuid"`
	Config          json.RawMessage `gorm:"type:jsonb"` // Pipeline-wide settings
	CreatedByUserID uuid.UUID       `gorm:"type:uuid"`
	UpdatedByUserID uuid.UUID       `gorm:"type:uuid"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Associations
	Source                  *Source                    `gorm:"foreignKey:TenantID,SourceID"`
	TransformerAssocs       []PipelineTransformer      `gorm:"foreignKey:TenantID,PipelineID"`
	DestinationInputsAssocs []PipelineDestinationInput `gorm:"foreignKey:TenantID,PipelineID"`
	Destination             Destination                `gorm:"foreignKey:TenantID,DestinationID"`
	Tenant                  Tenant                     `gorm:"foreignKey:TenantID"`
	CreatedBy               User                       `gorm:"foreignKey:CreatedByUserID"`
	UpdatedBy               User                       `gorm:"foreignKey:UpdatedByUserID"`
}

type PipelineTransformer struct {
	TenantID      uuid.UUID       `gorm:"primaryKey;type:uuid;uniqueIndex:idx_pipeline_stage"`
	PipelineID    uuid.UUID       `gorm:"primaryKey;type:uuid;uniqueIndex:idx_pipeline_stage"`
	TransformerID uuid.UUID       `gorm:"primaryKey;type:uuid"`
	Stage         int             `gorm:"type:int;uniqueIndex:idx_pipeline_stage"` // Order in pipeline (1, 2, 3, ...)
	Config        json.RawMessage `gorm:"type:jsonb"`                              // Transformer-specific pipeline config
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Associations
	Transformer Transformer                `gorm:"foreignKey:TenantID,TransformerID"`
	Inputs      []PipelineTransformerInput `gorm:"foreignKey:TenantID,PipelineID,TransformerID"`
}

type PipelineTransformerInput struct {
	TenantID      uuid.UUID `gorm:"primaryKey;type:uuid"`
	PipelineID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	TransformerID uuid.UUID `gorm:"primaryKey;type:uuid"`
	TopicID       uuid.UUID `gorm:"primaryKey;type:uuid"` // Input topic from previous stage
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Associations
	Topic Topic `gorm:"foreignKey:TenantID,TopicID"`
}

type PipelineDestinationInput struct {
	TenantID   uuid.UUID `gorm:"primaryKey;type:uuid"`
	PipelineID uuid.UUID `gorm:"primaryKey;type:uuid"`
	TopicID    uuid.UUID `gorm:"primaryKey;type:uuid"` // Output topics from last transformer or source to destination as input
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Associations
	Topic Topic `gorm:"foreignKey:TenantID,TopicID"`
}
