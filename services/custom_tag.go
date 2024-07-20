package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/custom_tag"

	"github.com/google/uuid"
)

type CustomTagService struct {
	repo repositories.CustomTagRepository
}

func NewCustomTagService(repo repositories.CustomTagRepository) *CustomTagService {
	return &CustomTagService{repo: repo}
}

func (s *CustomTagService) GetCustomTagsByUser(id string) (*[]interfaces.CustomTagResultQuery, error) {
	res, err := s.repo.GetCustomTagsByUser(id)
	return res, err
}

func (s *CustomTagService) CreateCustomTag(id uuid.UUID) (*uuid.UUID, error) {
	res, err := s.repo.CreateCustomTag(id, interfaces.CustomTagInsertDB{
		NameTh:         "รถเมย์",
		NameEn:         "bus",
		SpendingTypeId: uuid.MustParse("3ca5eecd-0a40-49f6-91bf-f5761c04e7f2"),
	})
	return res, err
}
