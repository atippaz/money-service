package repositories

import (
	"money-service/src/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userShareRepositoryGorm struct {
	db *gorm.DB
}

func NewGormUserShareRepository(db *gorm.DB) UserShareRepository {
	return &userShareRepositoryGorm{db: db}
}

func (r *userShareRepositoryGorm) GetAll(shareId *uuid.UUID) (*[]UserShareResultQuery, error) {
	var results []entities.UserShareEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	// var UserShareResults []UserShareResultQuery
	// for _, result := range results {
	// 	UserShareResults = append(UserShareResults, UserShareResultQuery{
	// 		UserShareId: result.UserShareId,
	// 		NameTh:      result.NameTh,
	// 		NameEn:      result.NameEn,
	// 	})
	// }

	// return &UserShareResults, nil
	return nil, nil

}
func (r *userShareRepositoryGorm) Insert(userOwner uuid.UUID, payload UserShareInsertDB) (*[]UserShareResultQuery, error) {
	var results []entities.UserShareEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	// var UserShareResults []UserShareResultQuery
	// for _, result := range results {
	// 	UserShareResults = append(UserShareResults, UserShareResultQuery{
	// 		UserShareId: result.UserShareId,
	// 		NameTh:      result.NameTh,
	// 		NameEn:      result.NameEn,
	// 	})
	// }

	// return &UserShareResults, nil
	return nil, nil

}
