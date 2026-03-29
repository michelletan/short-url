package config

import (
    "os"
    "strconv"
    "time"
)

type Config struct {
	DBURL    string
    JWTSecret string
    JWTTTL    time.Duration
}

func Load() (*Config, error) {
    ttlStr := os.Getenv("JWT_ACCESS_TOKEN_TTL")
    ttlSeconds, err := strconv.Atoi(ttlStr)
    if err != nil {
        return nil, err
    }

    return &Config{
		DBURL:    os.Getenv("DB_URL"),
        JWTSecret: os.Getenv("JWT_SECRET"),
        JWTTTL:    time.Duration(ttlSeconds) * time.Second,
    }, nil
}