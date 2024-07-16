package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IJwtHelper interface {
	CreateJWT(claims interface{}, expirationTime time.Time) (string, error)
	DecodeJWT(tokenStr string) (interface{}, error)
}

type jwtHelper struct {
}

var jwtKey = os.Getenv("SECRET_KEY")

func JwtHelper() IJwtHelper {
	return &jwtHelper{}
}

func (j *jwtHelper) CreateJWT(claims interface{}, expirationTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": claims,
		"exp":  expirationTime.Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *jwtHelper) DecodeJWT(tokenStr string) (interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if data, ok := claims["data"].(map[string]interface{}); ok {
			return data, nil
		}
	}

	return nil, errors.New("invalid token")
}
