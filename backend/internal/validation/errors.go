package validation

import "errors"

var (
    ErrPasswordLengthInvalid = errors.New("password must be between 8 - 30 characters")
    ErrUsernameLengthInvalid = errors.New("username must be between 8 - 30 characters")
	ErrEmailInvalid = errors.New("email format invalid")
)