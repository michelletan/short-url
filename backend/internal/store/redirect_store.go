package store

import (
    "database/sql"
    "short-url-backend/internal/models"
)

type RedirectStore struct {
    db *sql.DB
}

func NewRedirectStore(db *sql.DB) *RedirectStore {
	return &RedirectStore{db: db}
}

func (s *RedirectStore) Create(r *models.Redirect) error {
    query := `
        INSERT INTO redirects (url_id, user_ip, user_agent, referrer)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
    return s.db.QueryRow(query, r.URLID, r.UserIP, r.UserAgent, r.Referrer).
        Scan(&r.ID, &r.CreatedAt)
}