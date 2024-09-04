package services

import (
	share_repo "money-service/src/repositories/share"

	"github.com/google/uuid"
)

func NewShareService(repo share_repo.ShareRepository) ShareService {
	return &shareService{repo: repo}
}

func (s *shareService) GetById(shareId *uuid.UUID) (*[]share_repo.ShareResultQuery, error) {
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
func (s *shareService) GetAll(userId *uuid.UUID) (*[]share_repo.ShareResultQuery, error) {
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
func (s *shareService) Insert(userOwner uuid.UUID, payload share_repo.ShareInsertDB) (*[]share_repo.ShareResultQuery, error) {
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
