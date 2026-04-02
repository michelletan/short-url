package models

import "time"

type Link struct {
    ID         int
    UserID     int
    LongURL    string
    ShortCode  string
    ClickCount int
    CreatedAt  time.Time
    UpdatedAt  time.Time
}