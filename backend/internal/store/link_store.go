package store

import (
    "database/sql"
    "short-url-backend/internal/models"
)

type LinkStore struct {
    DB *sql.DB
}

func NewLinkStore(db *sql.DB) *LinkStore {
	return &LinkStore{DB: db}
}

func (s *LinkStore) Create(link *models.Link) error {
    query := `
        INSERT INTO links (user_id, long_url, short_code)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
    return s.DB.QueryRow(query, link.UserID, link.LongURL, link.ShortCode).
        Scan(&link.ID, &link.CreatedAt, &link.UpdatedAt)
}

func (s *LinkStore) GetByUserId(userID int) ([]models.Link, error) {
    query := `
        SELECT id, user_id, long_url, short_code, created_at, updated_at
        FROM links
        WHERE user_id = $1
        ORDER BY created_at DESC
    `

    rows, err := s.DB.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var links []models.Link
    for rows.Next() {
        var l models.Link
        if err := rows.Scan(
            &l.ID,
            &l.UserID,
            &l.LongURL,
            &l.ShortCode,
            &l.CreatedAt,
            &l.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        links = append(links, l)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return links, nil
}