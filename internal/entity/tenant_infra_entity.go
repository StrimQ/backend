package entity

import (
	"time"

	"github.com/google/uuid"
)

type TenantInfraEntity struct {
	TenantInfraID     uuid.UUID
	Name              string
	KafkaBrokers      string
	SchemaRegistryURL string
	KafkaConnectURL   string
	KmsKey            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
