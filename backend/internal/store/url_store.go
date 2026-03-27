package store

import (
    "database/sql"
    "short-url-backend/internal/models"
)

type URLStore struct {
    db *sql.DB
}

func NewURLStore(db *sql.DB) *URLStore {
	return &URLStore{db: db}
}

func (s *URLStore) Create(url *models.URL) error {
    query := `
        INSERT INTO urls (user_id, long_url, short_code)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
    return s.db.QueryRow(query, url.UserID, url.LongURL, url.ShortCode).
        Scan(&url.ID, &url.CreatedAt, &url.UpdatedAt)
}

func (s *URLStore) GetByShortCode(code string) (*models.URL, error) {
    url := &models.URL{}
    query := `
        SELECT id, user_id, long_url, short_code, created_at, updated_at
        FROM urls
        WHERE short_code = $1
    `
    row := s.db.QueryRow(query, code)
    if err := row.Scan(&url.ID, &url.UserID, &url.LongURL, &url.ShortCode, &url.CreatedAt, &url.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return url, nil
}