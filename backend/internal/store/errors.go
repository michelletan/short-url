package store

import "errors"

var (
	ErrDuplicateEmail = errors.New("Email already in use")
	ErrUserNotFound = errors.New("User not found")
)