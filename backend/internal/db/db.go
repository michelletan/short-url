package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
    "os"
)

func Connect() (*sql.DB, error) {
    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        return nil, fmt.Errorf("DB_URL not set")
    }

    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Optional: test connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return db, nil
}