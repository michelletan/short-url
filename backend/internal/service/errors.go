package service

import "errors"

var (
    ErrInternalServer = errors.New("Internal server error")
    ErrUserNotFound   = errors.New("User not found")
    ErrDuplicateEmail = errors.New("Email already in use")
    ErrInvalidLogin   = errors.New("Invalid email or password")
)