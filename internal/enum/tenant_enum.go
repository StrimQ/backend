package enum

type TenantTier string

const (
	TenantTier_FreeTrial TenantTier = "free_trial"
	TenantTier_Bronze    TenantTier = "bronze"
	TenantTier_Silver    TenantTier = "silver"
	TenantTier_Gold      TenantTier = "gold"
	TenantTier_Platinum  TenantTier = "platinum"
)
