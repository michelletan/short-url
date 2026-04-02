package store

import (
    "database/sql"
    "short-url-backend/internal/models"
)


type LinkStore interface {
    Create(link *models.Link) error
    GetByUserId(userID int) ([]*models.Link, error)
    GetByShortCode(shortCode string) (*models.Link, error)
    ExistsByShortCode(shortCode string) (bool, error)
    IncrementClickCount(linkID int) error
}

type LinkStoreImpl struct {
    DB *sql.DB
}

func NewLinkStore(db *sql.DB) *LinkStoreImpl {
	return &LinkStoreImpl{DB: db}
}

func (s *LinkStoreImpl) Create(link *models.Link) error {
    query := `
        INSERT INTO links (user_id, long_url, short_code)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
    return s.DB.QueryRow(query, link.UserID, link.LongURL, link.ShortCode).
        Scan(&link.ID, &link.CreatedAt, &link.UpdatedAt)
}

func (s *LinkStoreImpl) GetByUserId(userID int) ([]*models.Link, error) {
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

    var links []*models.Link
    for rows.Next() {
        l := &models.Link{}
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

func (s *LinkStoreImpl) GetByShortCode(shortCode string) (*models.Link, error) {
    query := `
        SELECT id, user_id, long_url, short_code, created_at, updated_at
        FROM links
        WHERE short_code = $1
    `

    var l models.Link
    err := s.DB.QueryRow(query, shortCode).Scan(
        &l.ID,
        &l.UserID,
        &l.LongURL,
        &l.ShortCode,
        &l.CreatedAt,
        &l.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }
    return &l, nil
}

func (s *LinkStoreImpl) ExistsByShortCode(shortCode string) (bool, error) {
    query := `SELECT COUNT(1) FROM links WHERE short_code = $1`
    var count int
    err := s.DB.QueryRow(query, shortCode).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (s *LinkStoreImpl) IncrementClickCount(linkID int) error {
    _, err := s.DB.Exec(
        `UPDATE links SET click_count = click_count + 1 WHERE id = $1`,
        linkID,
    )
    return err
}