package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

func NewJwt(JwtKey string) Jwt {
	return &_jwt{JwtKey}
}

func (j *_jwt) CreateJWT(claims *AuthClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.JwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *_jwt) DecodeJWT(tokenStr string) (*AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
