package store

import (
    "database/sql"

    "short-url-backend/internal/models"
)

type UserStore struct {
    db *sql.DB
}

func NewUserStore(database *sql.DB) (*UserStore) {
    return &UserStore{db: database}
}

func (s *UserStore) Create(user *models.User) error {
    query := `
        INSERT INTO users (username, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
    return s.db.QueryRow(query, user.Username, user.Email, user.PasswordHash).Scan(
        &user.ID, &user.CreatedAt, &user.UpdatedAt,
    )
}

func (s *UserStore) GetByEmail(email string) (*models.User, error) {
    user := &models.User{}
    query := `
        SELECT id, username, email, password_hash, created_at, updated_at
        FROM users
        WHERE email=$1
    `
    row := s.db.QueryRow(query, email)
    if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return user, nil
}

func (s *UserStore) GetByID(id int) (*models.User, error) {
    user := &models.User{}
    query := `
        SELECT id, username, email, password_hash, created_at, updated_at
        FROM users
        WHERE id=$1
    `
    row := s.db.QueryRow(query, id)
    if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return user, nil
}