package enum

type TopicProducerType string

const (
	TopicProducerType_Source      TopicProducerType = "source"
	TopicProducerType_Transformer TopicProducerType = "transformer"
)
