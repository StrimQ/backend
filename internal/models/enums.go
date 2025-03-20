package models

type Tier string

const (
	TierFreeTrial Tier = "free_trial" // Free Trial tier will get their infras data wiped after 30 days
	TierBronze    Tier = "bronze"
	TierSilver    Tier = "silver"
	TierGold      Tier = "gold"
	TierPlatinum  Tier = "platinum" // Platinum tier will get their own dedicated infra
)

// TopicProducerType indicates the type of producer for a topic
type TopicProducerType string

const (
	TopicProducerTypeSource      TopicProducerType = "source"
	TopicProducerTypeTransformer TopicProducerType = "transformer"
)

// SourceEngine specifies the engine type for a source
type SourceEngine string

const (
	SourceEngineMySQL      SourceEngine = "mysql"
	SourceEnginePostgreSQL SourceEngine = "postgresql"
)

// DestinationEngine specifies the engine type for a destination
type DestinationEngine string

const (
	DestinationEngineMySQL      DestinationEngine = "mysql"
	DestinationEnginePostgreSQL DestinationEngine = "postgresql"
)
