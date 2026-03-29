package dtos

type RegisterRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterResponse struct {
    Message string `json:"message"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"` // seconds until access token expires
}

type MeResponse struct {
	Email string `json:"email"`
	Username string `json:"username"`
}