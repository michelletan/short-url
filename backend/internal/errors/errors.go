package errors

import "errors"

var (
    ErrInternalServer = errors.New("Internal server error")
    ErrInvalidLogin   = errors.New("Invalid email or password")
    ErrNotFound       = errors.New("Resource not found")
)