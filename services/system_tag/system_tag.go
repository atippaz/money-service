package services

import (
	repositories "money-service/repositories/system_tag"
)

func NewSystemTagService(repo repositories.SystemTagRepository) SystemTagService {
	return &systemTagService{repo: repo}
}

func (s *systemTagService) GetAllSystemTags() (*[]SystemTagResult, error) {
	// res, err := s.repo.GetAllSystemTags()
	// return res, err
	return nil, nil
}
