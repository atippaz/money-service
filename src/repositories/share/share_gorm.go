package repositories

import (
	"money-service/src/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type shareRepositoryGorm struct {
	db *gorm.DB
}

func NewGormShareRepository(db *gorm.DB) ShareRepository {
	return &shareRepositoryGorm{db: db}
}

func (r *shareRepositoryGorm) GetAll(userId *uuid.UUID) (*[]ShareResultQuery, error) {
	var results []entities.ShareEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	// var ShareResults []ShareResultQuery
	// for _, result := range results {
	// 	ShareResults = append(ShareResults, ShareResultQuery{
	// 		ShareId: result.ShareId,
	// 		NameTh:  result.NameTh,
	// 		NameEn:  result.NameEn,
	// 	})
	// }

	// return &ShareResults, nil
	return nil, nil
}
func (r *shareRepositoryGorm) Insert(userOwner uuid.UUID, payload ShareInsertDB) (*[]ShareResultQuery, error) {
	var results []entities.ShareEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	// var ShareResults []ShareResultQuery
	// for _, result := range results {
	// 	ShareResults = append(ShareResults, ShareResultQuery{
	// 		ShareId: result.ShareId,
	// 		NameTh:  result.NameTh,
	// 		NameEn:  result.NameEn,
	// 	})
	// }

	// return &ShareResults, nil
	return nil, nil

}
func (r *shareRepositoryGorm) GetById(shareId *uuid.UUID) (*[]ShareResultQuery, error) {
	var results []entities.ShareEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	// var ShareResults []ShareResultQuery
	// for _, result := range results {
	// 	ShareResults = append(ShareResults, ShareResultQuery{
	// 		ShareId: result.ShareId,
	// 		NameTh:  result.NameTh,
	// 		NameEn:  result.NameEn,
	// 	})
	// }

	// return &ShareResults, nil
	return nil, nil

}
