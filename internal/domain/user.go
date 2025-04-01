package domain

import (
	"time"

	"github.com/google/uuid"
)

const ContextKey_User ContextKey = "user"

type User struct {
	TenantID  uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	// Associations
	Tenant Tenant
}

// NewUser creates a new User instance.
func NewUser(tenantID uuid.UUID, userID uuid.UUID) *User {
	return &User{
		TenantID: tenantID,
		UserID:   userID,
	}
}

type Tenant struct {
	TenantID  string
	Name      string
	Domain    string
	Tier      string
	InfraID   uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	// Associations
	Infra TenantInfra
}

func NewTenant(tenantID, name, domain, tier string, infra TenantInfra) *Tenant {
	return &Tenant{
		TenantID: tenantID,
		Name:     name,
		Domain:   domain,
		Tier:     tier,
		Infra:    infra,
	}
}

type TenantInfra struct {
	Name              string
	KafkaBrokers      []string
	SchemaRegistryURL string
	KafkaConnectURL   string
	KmsKey            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
