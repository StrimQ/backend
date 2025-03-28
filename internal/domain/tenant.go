package domain

type Tenant struct {
	TenantID string
	Name     string
	Domain   string
	Tier     string
	Infra    *TenantInfra
}

func NewTenant(tenantID, name, domain, tier string, infra *TenantInfra) *Tenant {
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
}
