package services

import (
	userRepositories "money-service/src/repositories/user"
	Hasher "money-service/src/utils/hasher"

	"github.com/google/uuid"
)

func NewUserService(repo userRepositories.UserRepository, encode Hasher.Hasher) UserService {
	return &userService{repo: repo, encode: encode}
}

func (s *userService) GetUserById(userId uuid.UUID) (*UserInfo, error) {
	result, err := s.repo.GetUserById(userId)
	return &UserInfo{
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

func (s *userService) GetLoginDataByCredential(credential string) (*UserLoginInfo, error) {
	result, err := s.repo.GetUserByCredential(credential)
	if err != nil {
		return nil, err
	}
	return &UserLoginInfo{
		UserName: result.UserName,
		Email:    result.Email,
		Password: result.Password,
		UserId:   result.UserId,
	}, err
}

func (s *userService) DeActiveAccount(userId uuid.UUID) (bool, error) {
	res, err := s.repo.DeActiveAccount(userId)
	return res, err
}

func (s *userService) CreateUser(payload UserResult) (*uuid.UUID, error) {
	//todo check email
	//todo check username
	//todo hash password
	password, err := s.encode.Hashing(&payload.Password)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	res, err := s.repo.CreateUser(userRepositories.UserInsertDb{
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
