package entity

import (
	"time"

	"github.com/google/uuid"
)

type TenantUserEntity struct {
	TenantID  uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
