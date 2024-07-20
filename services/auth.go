package services

import (
	"fmt"
	"money-service/interfaces"
	"money-service/utils"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type AuthService struct {
	userService *UserService
	hasher      utils.Hasher
	jwt         utils.Jwt
}

func NewAuthService(userService *UserService, hasher utils.Hasher, jwt utils.Jwt) *AuthService {
	return &AuthService{userService: userService, hasher: hasher, jwt: jwt}
}

var loginWhiteLists string

func (s *AuthService) Register(payload interfaces.AuthRegisterInsert) (*uuid.UUID, error) {
	result, err := s.userService.CreateUser(interfaces.UserInsertDb{
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

func (s *AuthService) Login(credential, password string) (*string, error) {
	userSecureData, err := s.userService.GetLoginDataByCredential(credential)
	if err != nil {
		return nil, err
	}
	fmt.Println(userSecureData)
	if !s.hasher.ComparePassword(password, userSecureData.Password) {

		return nil, err
	}

	claims := interfaces.AuthClaims{
		UserId:   userSecureData.UserId,
		Username: userSecureData.UserName,
		Email:    userSecureData.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	jwt, err := s.jwt.CreateJWT(&claims)

	if err != nil {
		return nil, err
	}
	return &jwt, nil
}
