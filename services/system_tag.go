package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/system_tag"
)

type SystemTagService struct {
	repo repositories.SystemTagRepository
}

func NewSystemTagService(repo repositories.SystemTagRepository) *SystemTagService {
	return &SystemTagService{repo: repo}
}

func (s *SystemTagService) GetAllSystemTags() (*[]interfaces.SystemTagResultQuery, error) {
	res, err := s.repo.GetAllSystemTags()
	return res, err
}
