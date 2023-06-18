package dto

import "github.com/golang-jwt/jwt/v5"

// This is like the input!
type LoginRequestBody struct {
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	ClientID     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Token     string `json:"access_token"`
	Type      string `json:"Bearer"`
	ExpiredAt string `json:"expired_at"`
	Scope     string `json:"scope"`
}
