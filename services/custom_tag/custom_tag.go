package services

import (
	repositories "money-service/repositories/custom_tag"

	"github.com/google/uuid"
)

func NewCustomTagService(repo repositories.CustomTagRepository) CustomTagService {
	return &customTagService{repo: repo}
}
func (s *customTagService) GetCustomTagsByUser(id string) (*[]CustomTagResult, error) {
	res, err := s.repo.GetCustomTagsByUser(id)
	var results []CustomTagResult
	for _, result := range *res {
		results = append(results, CustomTagResult{
			TagId:          result.TagId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			IsActive:       result.IsActive,
			SpendingTypeId: result.SpendingTypeId,
			UserOwner:      result.UserOwner,
		})
	}
	return &results, err
}

func (s *customTagService) CreateCustomTag(id uuid.UUID) (*uuid.UUID, error) {
	res, err := s.repo.CreateCustomTag(id, repositories.CustomTagInsertDB{
		NameTh:         "รถเมย์",
		NameEn:         "bus",
		SpendingTypeId: uuid.MustParse("3ca5eecd-0a40-49f6-91bf-f5761c04e7f2"),
	})
	return res, err
}
