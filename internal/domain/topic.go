package domain

import "github.com/google/uuid"

type Topic struct {
	TenantID uuid.UUID `validate:"required"`
	TopicID  uuid.UUID `validate:"required"`
	Name     string    `validate:"required"`
}
