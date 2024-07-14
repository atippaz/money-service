package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/system_tag"
)

type ISystemTagService struct {
	repo repositories.ISystemTagRepository
}

func SystemTagService(repo repositories.ISystemTagRepository) *ISystemTagService {
	return &ISystemTagService{repo: repo}
}

func (s *ISystemTagService) GetAllSystemTags() (*[]interfaces.SystemTagResultQuery, error) {
	res, err := s.repo.GetAllSystemTags()
	return res, err
}
