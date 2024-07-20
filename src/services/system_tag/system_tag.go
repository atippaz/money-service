package services

import (
	repositories "money-service/src/repositories/system_tag"
)

func NewSystemTagService(repo repositories.SystemTagRepository) SystemTagService {
	return &systemTagService{repo: repo}
}

func (s *systemTagService) GetAllSystemTags() (*[]SystemTagResult, error) {
	res, err := s.repo.GetAllSystemTags()
	var results []SystemTagResult
	for _, result := range *res {
		results = append(results, SystemTagResult{
			TagId:          result.TagId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			IsActive:       result.IsActive,
			SpendingTypeId: result.SpendingTypeId,
		})
	}
	return &results, err

}
