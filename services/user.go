package services

import (
	"fmt"
	"money-service/interfaces"
	repositories "money-service/repositories/user"
	"money-service/utils"

	"github.com/google/uuid"
)

type UserService struct {
	repo   repositories.UserRepository
	encode utils.Hasher
}

func NewUserService(repo repositories.UserRepository, encode utils.Hasher) *UserService {
	fmt.Println(encode)
	return &UserService{repo: repo, encode: encode}
}

func (s *UserService) GetUserById(id uuid.UUID) (*interfaces.UserResultResponse, error) {
	result, err := s.repo.GetUserById(id)
	return &interfaces.UserResultResponse{
		UserId:      result.UserId,
		UserName:    result.UserName,
		Email:       result.Email,
		LastName:    result.LastName,
		FirstName:   result.FirstName,
		DisplayName: result.DisplayName,
		CreatedDate: result.CreatedDate,
		Profile:     result.Profile,
	}, err
}

func (s *UserService) GetLoginDataByCredential(credential string) (*interfaces.UserLoginInfoQuery, error) {
	result, err := s.repo.GetUserByCredential(credential)
	return &interfaces.UserLoginInfoQuery{
		UserName: result.UserName,
		Email:    result.Email,
		Password: result.Password,
		UserId:   result.UserId,
	}, err
}

func (s *UserService) DeActiveAccount(id string) (bool, error) {
	res, err := s.repo.DeActiveAccount(id)
	return res, err
}

func (s *UserService) CreateUser(payload interfaces.UserInsertDb) (*uuid.UUID, error) {
	//todo check email
	//todo check username
	//todo hash password
	fmt.Println(s.encode)
	password, err := s.encode.Hashing(&payload.Password)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	res, err := s.repo.CreateUser(interfaces.UserInsertDb{
		UserName:    payload.UserName,
		Email:       payload.Email,
		LastName:    payload.LastName,
		FirstName:   payload.FirstName,
		DisplayName: payload.DisplayName,
		Password:    password,
		Profile:     payload.Profile,
	})

	return res, err
}
