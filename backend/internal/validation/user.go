package validation

import (
    "net/mail"
)

// ValidatePassword checks if password length is within bounds
func ValidatePassword(password string) error {
    if len(password) < 8 || len(password) > 30 {
        return ErrPasswordLengthInvalid
    }
    return nil
}

// ValidateUsername checks if username length is within bounds
func ValidateUsername(username string) error {
    if len(username) < 8 || len(username) > 30 {
        return ErrUsernameLengthInvalid
    }
    return nil
}

// ValidateEmail checks if email is valid format
func ValidateEmail(email string) error {
    _, err := mail.ParseAddress(email)
    if err != nil {
        return ErrEmailInvalid
    }
    return nil
}

func ValidateUser(username, email, password string) error {
	if err := ValidateUsername(username); err != nil {
        return err
    }

    if err := ValidateEmail(email); err != nil {
        return err
    }

    if err := ValidatePassword(password); err != nil {
        return err
    }

    return nil
}