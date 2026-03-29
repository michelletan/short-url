package db

import (
    "github.com/lib/pq"
)

func IsUniqueViolation(err error) bool {
    if pqErr, ok := err.(*pq.Error); ok {
        return pqErr.Code == "23505"
    }
    return false
}