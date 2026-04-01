package service

import (
    "log"

    "short-url-backend/internal/store"
    "short-url-backend/internal/models"
)


type RedirectStore interface {
    Create(r *models.Redirect) error
}

type RedirectService struct {
    LinkStore store.LinkStore
    RedirectStore RedirectStore
}

func NewRedirectService(linkStore store.LinkStore, redirectStore RedirectStore) *RedirectService {
    return &RedirectService{LinkStore: linkStore, RedirectStore: redirectStore}
}

func (s *RedirectService) GetLinkByShortCode(slug string) (*models.Link, error) {
    link, err := s.LinkStore.GetByShortCode(slug)
    if err != nil {
        log.Printf("Error fetching link for short code %s: %v", slug, err)
        return nil, err
    }
    return link, nil
}

func (s *RedirectService) TrackRedirect(urlID int, userIP, userAgent, referrer string) (error) {
    r := &models.Redirect{
        URLID:     urlID,
        UserIP:    userIP,
        UserAgent: userAgent,
        Referrer:  referrer,
    }
    if err := s.RedirectStore.Create(r); err != nil {
        log.Printf("Error tracking redirect for URL ID %d: %v", urlID, err)
        return err
    }
    return nil
}