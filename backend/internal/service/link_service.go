package service

import (
    "fmt"
    "log"

    "short-url-backend/internal/store"
    "short-url-backend/internal/models"
    "short-url-backend/internal/util"
)

type LinkService struct {
    store store.LinkStore
}

func NewLinkService(store store.LinkStore) *LinkService {
    return &LinkService{store: store}
}

// CreateShortLink generates a unique short code and saves the URL
func (s *LinkService) CreateShortLink(userID int, originalURL string) (*models.Link, error) {
    code, err := generateUniqueShortCode(6, s.store)
    if err != nil {
        log.Printf("Error generating unique short code for user %d: %v", userID, err)
        return nil, err
    }

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

func generateUniqueShortCode(length int, store store.LinkStore) (string, error) {
    const maxAttempts = 10

    for i := 0; i < maxAttempts; i++ {
        code, err := util.GenerateShortCode(length)
        if err != nil {
            return "", err
        }

        exists, err := store.ExistsByShortCode(code)
        if err != nil {
            return "", err
        }
        if !exists {
            return code, nil
        }
    }

    return "", fmt.Errorf("could not generate unique code after %d attempts", maxAttempts)
}