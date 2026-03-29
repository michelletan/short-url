package service

import (
    "log"
    "errors"

    "golang.org/x/crypto/bcrypt"

    "short-url-backend/internal/models"
    "short-url-backend/internal/dtos"
    "short-url-backend/internal/store"
)

type UserStore interface {
    Create(user *models.User) error
    GetByEmail(email string) (*models.User, error)
    GetByID(id int) (*models.User, error)
}

type UserService struct {
    store UserStore
    jwtService *JWTService
}

func NewUserService(store UserStore, jwtService *JWTService) *UserService {
    return &UserService{store: store, jwtService: jwtService}
}

func (s *UserService) Register(username, email, password string) (*models.User, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Error hashing password for user %s: %v", email, err)
        return nil, ErrInternalServer
    }

    user := &models.User{
        Username:     username,
        Email:        email,
        PasswordHash: string(hashed),
    }

    if err := s.store.Create(user); err != nil {
        if errors.Is(err, store.ErrDuplicateEmail) {
            log.Printf("Duplicate email registration attempt for %s", email)
            return nil, ErrDuplicateEmail
        }
        log.Printf("Error creating user %s: %v", email, err)
        return nil, ErrInternalServer
    }

    return user, nil
}

func (s *UserService) Login(email, password string) (dtos.LoginResponse, error) {
    user, err := s.store.GetByEmail(email)
    if err != nil {
        if errors.Is(err, store.ErrUserNotFound) {
            log.Printf("No such user: %v", err)
            return dtos.LoginResponse{}, ErrInvalidLogin
        }
        log.Printf("Error fetching user %s: %v", email, err)
        return dtos.LoginResponse{}, ErrInvalidLogin
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        log.Printf("Password mismatch for user %s: %v", email, err)
        return dtos.LoginResponse{}, ErrInvalidLogin
    }

    access, err := s.jwtService.GenerateAccessToken(user.ID)
    if err != nil {
        log.Printf("Error generating access token for user %s: %v", email, err)
        return dtos.LoginResponse{}, ErrInternalServer
    }

    return dtos.LoginResponse{
        AccessToken:  access,
        ExpiresIn:    int(s.jwtService.accessTokenTTL.Seconds()),
    }, nil
}

func (s *UserService) GetByID(id int) (*models.User, error) {
    return s.store.GetByID(id)
}