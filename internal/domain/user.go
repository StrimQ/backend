package domain

import "github.com/google/uuid"

type User struct {
	TenantID uuid.UUID
	UserID   uuid.UUID
}
