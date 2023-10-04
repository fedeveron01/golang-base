package core_errors

type BadRequestError struct {
	Message string
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{Message: message}
}

func (b BadRequestError) Error() string {
	return b.Message
}
