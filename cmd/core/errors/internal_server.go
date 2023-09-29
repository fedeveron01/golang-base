package core_errors

type InternalServerError struct {
	Message string
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{Message: message}
}

func (i *InternalServerError) Error() string {
	return i.Message
}
