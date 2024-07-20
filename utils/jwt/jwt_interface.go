package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Jwt interface {
	CreateJWT(claims *AuthClaims) (string, error)
	DecodeJWT(tokenStr string) (*AuthClaims, error)
}

type _jwt struct {
	JwtKey string
}
type AuthClaims struct {
	Username string    `json:"userName"`
	Email    string    `json:"email"`
	UserId   uuid.UUID `json:"userId"`
	jwt.RegisteredClaims
}
