package models

import "time"

type Redirect struct {
    ID        int
    URLID     int
    UserIP    string
    UserAgent string
    Referrer  string
    CreatedAt time.Time
}