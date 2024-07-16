package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/user"

	"github.com/google/uuid"
)

type IUserService struct {
	repo repositories.IUserRepository
}

func UserService(repo repositories.IUserRepository) *IUserService {
	return &IUserService{repo: repo}
}

func (s *IUserService) GetUserById(id uuid.UUID) (*interfaces.UserResultResponse, error) {
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

func (s *IUserService) GetLoginDataByCredential(credential string) (*interfaces.UserLoginInfoQuery, error) {
	result, err := s.repo.GetUserByCredential(credential)
	return &interfaces.UserLoginInfoQuery{
		UserName: result.UserName,
		Email:    result.Email,
		Password: result.Password,
		UserId:   result.UserId,
	}, err
}

func (s *IUserService) DeActiveAccount(id string) (bool, error) {
	res, err := s.repo.DeActiveAccount(id)
	return res, err
}

func (s *IUserService) CreateUser(payload interfaces.UserInsertDb) (*uuid.UUID, error) {
	//todo check email
	//todo check username
	//todo hash password

	res, err := s.repo.CreateUser(interfaces.UserInsertDb{
		UserName:    payload.UserName,
		Email:       payload.Email,
		LastName:    payload.LastName,
		FirstName:   payload.FirstName,
		DisplayName: payload.DisplayName,
		Password:    payload.Password,
		Profile:     payload.Profile,
	})
	return res, err
}
