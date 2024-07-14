package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/user"
)

type IUserService struct {
	repo repositories.IUserRepository
}

func UserService(repo repositories.IUserRepository) *IUserService {
	return &IUserService{repo: repo}
}

func (s *IUserService) GetUserById(id string) (*interfaces.UserResultQuery, error) {
	res, err := s.repo.GetUserById(id)
	return res, err
}
