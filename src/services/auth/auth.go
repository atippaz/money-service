package services

import (
	"fmt"
	user "money-service/src/services/user"
	Hasher "money-service/src/utils/hasher"
	Jwt "money-service/src/utils/jwt"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func NewAuthService(userService user.UserService, hasher Hasher.Hasher, _jwt Jwt.Jwt) AuthService {
	return &authService{userService: userService, hasher: hasher, jwt: _jwt}
}

// var loginWhiteLists string

func (s *authService) Register(payload AuthRegisterInsert) (*uuid.UUID, error) {
	result, err := s.userService.CreateUser(user.UserResult{
		UserName:    payload.UserName,
		Email:       payload.Email,
		LastName:    payload.LastName,
		FirstName:   payload.FirstName,
		DisplayName: payload.DisplayName,
		Password:    payload.Password,
		Profile:     payload.Profile,
	})
	return result, err
}

func (s *authService) Login(credential, password string) (*string, error) {
	userSecureData, err := s.userService.GetLoginDataByCredential(credential)
	if err != nil {
		return nil, err
	}
	fmt.Println(userSecureData)
	if !s.hasher.ComparePassword(password, userSecureData.Password) {
		return nil, err
	}

	claims := AuthClaims{
		UserId:   userSecureData.UserId,
		Username: userSecureData.UserName,
		Email:    userSecureData.Email,
	}
	jwt, err := s.jwt.CreateJWT(&Jwt.AuthClaims{
		Username: claims.Username, Email: claims.Email, UserId: claims.UserId, RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	if err != nil {
		return nil, err
	}
	return &jwt, nil
}
