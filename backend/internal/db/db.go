package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func Connect(dbURL string) (*sql.DB, error) {
    if dbURL == "" {
        return nil, fmt.Errorf("DB_URL not set")
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        return nil, fmt.Errorf("Failed to open database: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("Failed to ping database: %w", err)
    }

    return db, nil
}