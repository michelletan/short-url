package service

import "errors"

var (
    ErrInternalServer = errors.New("internal server error")
    ErrUserNotFound   = errors.New("user not found")
    ErrInvalidUserForm = errors.New("invalid user form")
    ErrDuplicateEmail = errors.New("email already in use")
    ErrInvalidLogin   = errors.New("invalid email or password")
)