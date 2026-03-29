package service

import (
    "math/rand"
    "time"

    "short-url-backend/internal/models"
    "short-url-backend/internal/store"
)

type LinkService struct {
    store *store.LinkStore
}

func NewLinkService(store *store.LinkStore) *LinkService {
    return &LinkService{store: store}
}

// CreateShortLink generates a unique short code and saves the URL
func (s *LinkService) CreateShortLink(userID int, longURL string) (*models.Link, error) {
    code := generateShortCode(6)
    link := &models.Link{
        UserID:    userID,
        LongURL:   longURL,
        ShortCode: code,
    }
    if err := s.store.Create(link); err != nil {
        return nil, err
    }
    return link, nil
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