package service

import (
    "log"

    "short-url-backend/internal/store"
    "short-url-backend/internal/models"
)


type RedirectEventStore interface {
    Create(r *models.RedirectEvent) error
}

type RedirectService struct {
    LinkStore store.LinkStore
    RedirectEventStore RedirectEventStore
}

func NewRedirectService(linkStore store.LinkStore, redirectEventStore RedirectEventStore) *RedirectService {
    return &RedirectService{LinkStore: linkStore, RedirectEventStore: redirectEventStore}
}

func (s *RedirectService) GetLinkByShortCode(slug string) (*models.Link, error) {
    link, err := s.LinkStore.GetByShortCode(slug)
    if err != nil {
        log.Printf("Error fetching link for short code %s: %v", slug, err)
        return nil, err
    }
    return link, nil
}

func (s *RedirectService) TrackRedirect(urlID int, userIP, userAgent, referrer string) error {
    r := &models.RedirectEvent{
        URLID:     urlID,
        UserIP:    userIP,
        UserAgent: userAgent,
        Referrer:  referrer,
    }
    if err := s.RedirectEventStore.Create(r); err != nil {
        log.Printf("Error tracking redirect for URL ID %d: %v", urlID, err)
        return err
    }

    // increment cached counter on the link
    if err := s.LinkStore.IncrementClickCount(urlID); err != nil {
        // non-fatal — redirect event is already recorded
        log.Printf("Error incrementing click count for URL ID %d: %v", urlID, err)
    }

    return nil
}