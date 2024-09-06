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

	if err := db.Select("user_share_id, user_id, share_id").Find(&results).Error; err != nil {
		return nil, err
	}
	var UserShareResults []UserShareResultQuery
	for _, result := range results {
		UserShareResults = append(UserShareResults, UserShareResultQuery{
			UserShareId: result.UserShareId,
			UserId:      result.UserId,
			ShareId:     result.ShareId,
		})
	}

	return &UserShareResults, nil

}
func (r *userShareRepositoryGorm) Insert(userOwner uuid.UUID, payload UserShareInsertDB) (*uuid.UUID, error) {
	db := r.db
	newTag := entities.UserShareEntity{
		UserId:  payload.UserId,
		ShareId: userOwner,
	}
	if err := db.Create(&newTag).Error; err != nil {
		return nil, err
	}
	return &newTag.UserShareId, nil
}
