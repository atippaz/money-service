package services

import (
	user_share_repo "money-service/src/repositories/user_share"

	"github.com/google/uuid"
)

func NewUserShareService(repo user_share_repo.UserShareRepository) UserShareService {
	return &userShareService{repo: repo}
}

func (s *userShareService) GetAll(userId *uuid.UUID) (*[]user_share_repo.UserShareResultQuery, error) {
	// res, err := s.repo.GetShares()
	// var results []ShareResult
	// for _, data := range *res {
	// 	results = append(results, ShareResult{
	// 		ShareId: data.ShareId,
	// 		NameTh:  data.NameTh,
	// 		NameEn:  data.NameEn,
	// 	})
	// }
	// return &results, err
	return nil, nil

}
func (s *userShareService) Insert(userOwner uuid.UUID, payload user_share_repo.UserShareInsertDB) (*[]user_share_repo.UserShareResultQuery, error) {
	// res, err := s.repo.GetShares()
	// var results []ShareResult
	// for _, data := range *res {
	// 	results = append(results, ShareResult{
	// 		ShareId: data.ShareId,
	// 		NameTh:  data.NameTh,
	// 		NameEn:  data.NameEn,
	// 	})
	// }
	// return &results, err
	return nil, nil

}
