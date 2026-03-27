package service

import (
    "short-url-backend/internal/models"
    "short-url-backend/internal/store"
    "math/rand"
    "time"
)

type URLService struct {
    store *store.URLStore
}

func NewURLService(store *store.URLStore) *URLService {
    return &URLService{store: store}
}

// CreateShortURL generates a unique short code and saves the URL
func (s *URLService) CreateShortURL(userID int, longURL string) (*models.URL, error) {
    code := generateShortCode(6)
    url := &models.URL{
        UserID:    userID,
        LongURL:   longURL,
        ShortCode: code,
    }
    if err := s.store.Create(url); err != nil {
        return nil, err
    }
    return url, nil
}

// Generates a random alphanumeric short code
func generateShortCode(length int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    rand.Seed(time.Now().UnixNano())
    b := make([]byte, length)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}