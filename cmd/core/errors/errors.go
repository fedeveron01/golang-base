package core_errors

import "errors"

var (
	ErrInvalidUsernameOrPassword     = errors.New("invalid username or password")
	ErrUserIsInactive                = errors.New("user is inactive, please contact with the administrator")
	ErrUserNotFound                  = errors.New("user not found")
	ErrUsernameOrPasswordIsIncorrect = errors.New("username or password is incorrect")
	ErrUsernameOrPasswordIsEmpty     = errors.New("username or password is empty")
	ErrUsernameAlreadyExists         = errors.New("username already exists")
	ErrInactiveUser                  = errors.New("user is inactive, please contact with the administrator")
)
