package services

import (
	user_share_repo "money-service/src/repositories/user_share"

	"github.com/google/uuid"
)

func NewUserShareService(repo user_share_repo.UserShareRepository) UserShareService {
	return &userShareService{repo: repo}
}

func (s *userShareService) GetAll(userId *uuid.UUID) (*[]user_share_repo.UserShareResultQuery, error) {
	res, err := s.repo.GetAll(userId)
	var results []user_share_repo.UserShareResultQuery
	for _, data := range *res {
		results = append(results, user_share_repo.UserShareResultQuery{
			ShareId:     data.ShareId,
			UserShareId: data.UserShareId,
			UserId:      data.UserId,
		})
	}
	return &results, err

}
func (s *userShareService) Insert(userOwner uuid.UUID, payload user_share_repo.UserShareInsertDB) (*uuid.UUID, error) {
	res, err := s.repo.Insert(userOwner, user_share_repo.UserShareInsertDB{
		UserId: payload.UserId,
	})
	return res, err
}
