package domain

import (
	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

type Topic struct {
	TenantID     uuid.UUID `validate:"required"`
	TopicID      uuid.UUID `validate:"required"`
	Name         string    `validate:"required"`
	ProducerType enum.TopicProducerType
	ProducerID   uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewTopic(
	tenantID uuid.UUID,
	topicID uuid.UUID,
	name string,
	producerType enum.TopicProducerType,
	producerID uuid.UUID,
) *Topic {
	return &Topic{
		TenantID:     tenantID,
		TopicID:      topicID,
		Name:         name,
		ProducerType: producerType,
		ProducerID:   producerID,
	}
}
