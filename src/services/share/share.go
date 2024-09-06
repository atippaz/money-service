package services

import (
	share_repo "money-service/src/repositories/share"

	"github.com/google/uuid"
)

func NewShareService(repo share_repo.ShareRepository) ShareService {
	return &shareService{repo: repo}
}

func (s *shareService) GetById(shareId *uuid.UUID) (*[]share_repo.ShareResultQuery, error) {
	res, err := s.repo.GetById(shareId)
	var results []share_repo.ShareResultQuery
	for _, data := range *res {
		results = append(results, share_repo.ShareResultQuery{
			ShareId:     data.ShareId,
			StartDate:   data.StartDate,
			EndDate:     data.EndDate,
			ExpiredDate: data.ExpiredDate,
			UserShareId: data.UserShareId,
		})
	}
	return &results, err
}
func (s *shareService) GetAll(userId *uuid.UUID) (*[]share_repo.ShareResultQuery, error) {
	res, err := s.repo.GetAll(userId)
	var results []share_repo.ShareResultQuery
	for _, data := range *res {
		results = append(results, share_repo.ShareResultQuery{
			ShareId:     data.ShareId,
			StartDate:   data.StartDate,
			EndDate:     data.EndDate,
			ExpiredDate: data.ExpiredDate,
			UserShareId: data.UserShareId,
		})
	}
	return &results, err
}
func (s *shareService) Insert(userOwner uuid.UUID, payload share_repo.ShareInsertDB) (*uuid.UUID, error) {
	res, err := s.repo.Insert(userOwner, share_repo.ShareInsertDB{
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		ExpiredDate: payload.ExpiredDate,
		UserShareId: userOwner,
	})
	return res, err

}
