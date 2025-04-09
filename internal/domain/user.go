package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	TenantID  uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time

	// Associations
	Tenant *Tenant
}

// NewUser creates a new User instance.
func NewUser(tenantID uuid.UUID, userID uuid.UUID) *User {
	return &User{
		TenantID: tenantID,
		UserID:   userID,
	}
}

type Tenant struct {
	TenantID      string
	Name          string
	Domain        string
	Tier          string
	TenantInfraID uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Associations
	Infra *TenantInfra
}

func NewTenant(tenantID, name, domain, tier string) *Tenant {
	return &Tenant{
		TenantID: tenantID,
		Name:     name,
		Domain:   domain,
		Tier:     tier,
	}
}

type TenantInfra struct {
	TenantInfraID     uuid.UUID
	Name              string
	KafkaBrokers      []string
	SchemaRegistryURL string
	KafkaConnectURL   string
	KmsKey            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
