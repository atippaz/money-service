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

func (s *ICustomTagService) GetCustomTagsByUser(id string) (*[]interfaces.CustomTagResultQuery, error) {
	res, err := s.repo.GetCustomTagsByUser(id)
	return res, err
}

// func (s *ICustomTagService) CreateCustomTag(id string) (*[]interfaces.CustomTagResultQuery, error) {
// 	res, err := s.repo.GetCustomTagsByUser(id)
// 	return res, err
// }
