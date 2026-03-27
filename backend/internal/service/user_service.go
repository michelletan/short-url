package service

import (
    "short-url-backend/internal/models"
    "short-url-backend/internal/store"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    store *store.UserStore
}

func NewUserService(store *store.UserStore) *UserService {
    return &UserService{store: store}
}

func (s *UserService) Register(username, email, password string) (*models.User, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Username:     username,
        Email:        email,
        PasswordHash: string(hashed),
    }

    if err := s.store.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}