package services

import (
	"money-service/interfaces"
	"money-service/utils"
	"time"

	"github.com/google/uuid"
)

type IAuthService struct {
	userService  *IUserService
	encodeHelper utils.IEncodingHelper
	jwtHelper    utils.IJwtHelper
}

func AuthService(userService *IUserService, encodeHelper utils.IEncodingHelper, jwtHelper utils.IJwtHelper) *IAuthService {
	return &IAuthService{userService: userService, encodeHelper: encodeHelper, jwtHelper: jwtHelper}
}

var loginWhiteLists string

func (s *IAuthService) Register(payload interfaces.AuthRegisterInsert) (*uuid.UUID, error) {
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

func (s *IAuthService) Login(credential, password string) (*string, error) {
	userSecureData, err := s.userService.GetLoginDataByCredential(credential)
	if err != nil {
		return nil, err
	}
	if !s.encodeHelper.ComparePassword(password, userSecureData.Password) {
		return nil, err
	}
	jwt, err := s.jwtHelper.CreateJWT(&interfaces.AuthClaims{
		UserId:   userSecureData.UserId,
		Username: userSecureData.UserName,
		Email:    userSecureData.Email,
	}, <-time.After(1))
	if err != nil {
		return nil, err
	}
	return &jwt, nil
}
