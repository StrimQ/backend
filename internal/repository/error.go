package repository

type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	return "user not found"
}
