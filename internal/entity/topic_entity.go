package entity

import (
	"time"

	"github.com/StrimQ/backend/internal/enum"
	"github.com/google/uuid"
)

type TopicEntity struct {
	TenantID     uuid.UUID
	TopicID      uuid.UUID
	Name         string
	ProducerType enum.TopicProducerType
	ProducerID   uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
