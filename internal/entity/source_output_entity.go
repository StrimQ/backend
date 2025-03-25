package entity

import (
	"time"

	"github.com/google/uuid"
)

type SourceOutputEntity struct {
	TenantID       uuid.UUID
	SourceID       uuid.UUID
	TopicID        uuid.UUID
	DatabaseName   string
	GroupName      string
	CollectionName string
	Config         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
