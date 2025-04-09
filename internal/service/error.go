package service

// ErrKCClientNotFound is a custom error for when the Kafka Connect client is not found in the context.
type ErrKCClientNotFound struct{}

func (e *ErrKCClientNotFound) Error() string {
	return "Kafka Connect client not found in context"
}
