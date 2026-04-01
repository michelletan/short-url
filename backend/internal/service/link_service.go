package service

import (
    "log"
    "math/rand"
    "time"

    "short-url-backend/internal/models"
)

type LinkStore interface {
    Create(link *models.Link) error
    GetByUserId(userID int) ([]models.Link, error)
}

type LinkService struct {
    store LinkStore
}

func NewLinkService(store LinkStore) *LinkService {
    return &LinkService{store: store}
}

// CreateShortLink generates a unique short code and saves the URL
func (s *LinkService) CreateShortLink(userID int, originalURL string) (*models.Link, error) {
    code := generateShortCode(6)
    link := &models.Link{
        UserID:    userID,
        LongURL:   originalURL,
        ShortCode: code,
    }
    if err := s.store.Create(link); err != nil {
        log.Printf("Error creating link for user %d: %v", userID, err)
        return nil, err
    }
    return link, nil
}

func (s *LinkService) GetUserLinks(userID int) ([]models.Link, error) {
    links, err := s.store.GetByUserId(userID)
    if err != nil {
        log.Printf("Error getting links for user %d: %v", userID, err)
        return nil, err
    }
    return links, nil
}

// Generates a random alphanumeric short code
func generateShortCode(length int) string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    b := make([]byte, length)
    for i := range b {
        b[i] = letters[r.Intn(len(letters))]
    }
    return string(b)
}