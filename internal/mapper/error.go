package mapper

// ErrUserNotFound is a custom error for when the user is not found in the context.
type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	return "user not found in context"
}
