package models

import (
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	TenantID  uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"type:varchar(255)"`
	Domain    string    `gorm:"type:varchar(255)"`
	Tier      Tier      `gorm:"type:varchar(255)"`
	InfraID   uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Associations
	Infra TenantInfra `gorm:"foreignKey:InfraID"`
}
type TenantUser struct {
	TenantID  uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Associations
	Tenant Tenant `gorm:"foreignKey:TenantID"`
	User   User   `gorm:"foreignKey:UserID"`
}

type TenantInfra struct {
	TenantInfraID     uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name              string    `gorm:"type:varchar(255)"`
	KafkaBrokers      []string  `gorm:"type:varchar(255)[]"`
	SchemaRegistryURL string    `gorm:"type:varchar(255)"`
	KafkaConnectURL   string    `gorm:"type:varchar(255)"`
	KMSKey            string    `gorm:"type:varchar(255)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
