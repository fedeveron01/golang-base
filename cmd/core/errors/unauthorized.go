package core_errors

type UnauthorizedError struct {
	Message string
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{Message: message}
}

func (u UnauthorizedError) Error() string {
	return u.Message
}
