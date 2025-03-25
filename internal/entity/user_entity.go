package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
