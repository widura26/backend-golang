package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Valid implements jwt.Claims.
func (j *JwtCustomClaims) Valid() error {
	panic("unimplemented")
}
