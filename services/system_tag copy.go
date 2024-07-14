package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/custom_tag"
)

type ICustomTagService struct {
	repo repositories.ICustomTagRepository
}

func CustomTagService(repo repositories.ICustomTagRepository) *ICustomTagService {
	return &ICustomTagService{repo: repo}
}

func (s *ICustomTagService) GetUserById(id string) (*interfaces.UserResult, error) {
	res, err := s.repo.GetUserById(id)
	return res, err
}
