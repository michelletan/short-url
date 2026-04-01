package store

import (
    "database/sql"
    "short-url-backend/internal/models"
)

type RedirectEventStore struct {
    db *sql.DB
}

func NewRedirectEventStore(db *sql.DB) *RedirectEventStore {
	return &RedirectEventStore{db: db}
}

func (s *RedirectEventStore) Create(r *models.RedirectEvent) error {
    query := `
        INSERT INTO redirect_events (url_id, user_ip, user_agent, referrer)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at
    `
    return s.db.QueryRow(query, r.URLID, r.UserIP, r.UserAgent, r.Referrer).
        Scan(&r.ID, &r.CreatedAt)
}