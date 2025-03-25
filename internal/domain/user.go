package domain

import "github.com/google/uuid"

const ContextKey_User ContextKey = "user"

type User struct {
	TenantID uuid.UUID
	UserID   uuid.UUID
}

// NewUser creates a new User instance.
func NewUser(tenantID uuid.UUID, userID uuid.UUID) *User {
	return &User{
		TenantID: tenantID,
		UserID:   userID,
	}
}
