package dtos

import "time"

type CreateLinkRequest struct {
    URL string `json:"url"`
}

type CreateLinkResponse struct {
    Link LinkDTO `json:"link"`
}

type GetLinksResponse struct {
    Links []LinkDTO `json:"links"`
}

type LinkDTO struct {
    ID         int    `json:"id"`
    URL        string    `json:"url"`
    Slug       string    `json:"slug"`
    CreatedAt  time.Time `json:"created_at"`
    ClickCount int       `json:"click_count"`
}