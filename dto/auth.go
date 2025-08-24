package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
