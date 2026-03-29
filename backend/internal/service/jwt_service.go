package service

import (
    "fmt"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
    secretKey       string
    accessTokenTTL  time.Duration
}

func NewJWTService(secret string, accessTTL time.Duration) *JWTService {
    return &JWTService{
        secretKey:       secret,
        accessTokenTTL:  accessTTL,
    }
}

// GenerateAccessToken returns a signed JWT string
func (s *JWTService) GenerateAccessToken(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(s.accessTokenTTL).Unix(),
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(s.secretKey))
}

// ValidateToken parses & validates a JWT
func (s *JWTService) ValidateToken(tokenStr string) (*jwt.Token, error) {
    return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(s.secretKey), nil
    })
}