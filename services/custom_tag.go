package services

import (
	"money-service/interfaces"
	repositories "money-service/repositories/custom_tag"

	"github.com/google/uuid"
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

func (s *ICustomTagService) CreateCustomTag(id uuid.UUID) (*uuid.UUID, error) {
	res, err := s.repo.CreateCustomTag(id, interfaces.CustomTagInsertDB{
		NameTh:         "รถเมย์",
		NameEn:         "bus",
		SpendingTypeId: uuid.MustParse("3ca5eecd-0a40-49f6-91bf-f5761c04e7f2"),
	})
	return res, err
}
