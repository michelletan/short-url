package models

import "time"

type RedirectEvent struct {
    ID        int
    URLID     int
    UserIP    string
    UserAgent string
    Referrer  string
    CreatedAt time.Time
}