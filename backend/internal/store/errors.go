package store

import "errors"

var (
	ErrDuplicateEmail = errors.New("email already in use")
	ErrUserNotFound = errors.New("user not found")
)