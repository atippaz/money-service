package services

import (
	repositories "money-service/src/repositories/tag"

	"github.com/google/uuid"
)

func NewTagService(repo repositories.TagRepository) TagService {
	return &tagService{repo: repo}
}
func (s *tagService) GetTagsByUser(userId uuid.UUID, hasSystem bool) (*[]TagResult, error) {
	res, err := s.repo.GetTags(&userId, hasSystem)
	var results []TagResult
	for _, result := range *res {
		results = append(results, TagResult{
			TagId:          result.TagId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			IsActive:       result.IsActive,
			SpendingTypeId: result.SpendingTypeId,
			Owner:          result.Owner,
		})
	}
	return &results, err
}
func (s *tagService) GetSystemTags() (*[]TagResult, error) {
	res, err := s.repo.GetTags(nil, true)
	var results []TagResult
	for _, result := range *res {
		results = append(results, TagResult{
			TagId:          result.TagId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			IsActive:       result.IsActive,
			SpendingTypeId: result.SpendingTypeId,
			Owner:          result.Owner,
		})
	}
	return &results, err
}

func (s *tagService) CreateTag(id uuid.UUID) (*uuid.UUID, error) {
	res, err := s.repo.CreateTag(id, repositories.TagInsertDB{
		NameTh:         "รถเมย์",
		NameEn:         "bus",
		SpendingTypeId: uuid.MustParse("3ca5eecd-0a40-49f6-91bf-f5761c04e7f2"),
	})
	return res, err
}
