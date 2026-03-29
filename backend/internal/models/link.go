package models

import "time"

type Link struct {
    ID        int
    UserID    int
    LongURL   string
    ShortCode string
    CreatedAt time.Time
    UpdatedAt time.Time
}