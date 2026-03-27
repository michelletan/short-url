package service

import (
    "short-url-backend/internal/models"
    "short-url-backend/internal/store"
)

type RedirectService struct {
    store *store.RedirectStore
}

func NewRedirectService(store *store.RedirectStore) *RedirectService {
    return &RedirectService{store: store}
}

func (s *RedirectService) TrackRedirect(urlID int, userIP, userAgent, referrer string) (*models.Redirect, error) {
    r := &models.Redirect{
        URLID:     urlID,
        UserIP:    userIP,
        UserAgent: userAgent,
        Referrer:  referrer,
    }
    if err := s.store.Create(r); err != nil {
        return nil, err
    }
    return r, nil
}