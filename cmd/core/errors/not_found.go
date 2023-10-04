package core_errors

type NotFoundError struct {
	Message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{Message: message}
}

func (n NotFoundError) Error() string {
	return n.Message
}
