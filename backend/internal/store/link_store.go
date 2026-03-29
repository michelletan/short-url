package store

import (
    "database/sql"
    "short-url-backend/internal/models"
)

type LinkStore struct {
    db *sql.DB
}

func NewLinkStore(db *sql.DB) *LinkStore {
	return &LinkStore{db: db}
}

func (s *LinkStore) Create(link *models.Link) error {
    query := `
        INSERT INTO urls (user_id, long_url, short_code)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
    return s.db.QueryRow(query, link.UserID, link.LongURL, link.ShortCode).
        Scan(&link.ID, &link.CreatedAt, &link.UpdatedAt)
}

func (s *LinkStore) GetByShortCode(code string) (*models.Link, error) {
    link := &models.Link{}
    query := `
        SELECT id, user_id, long_url, short_code, created_at, updated_at
        FROM urls
        WHERE short_code = $1
    `
    row := s.db.QueryRow(query, code)
    if err := row.Scan(&link.ID, &link.UserID, &link.LongURL, &link.ShortCode, &link.CreatedAt, &link.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return link, nil
}