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

	if err := db.Select("share_id, start_date, end_date,expired_date,user_share_id").Find(&results).Error; err != nil {
		return nil, err
	}
	var ShareResults []ShareResultQuery
	for _, result := range results {
		ShareResults = append(ShareResults, ShareResultQuery{
			ShareId:     result.ShareId,
			StartDate:   result.StartDate,
			EndDate:     result.EndDate,
			ExpiredDate: result.ExpiredDate,
			UserShareId: result.UserShareId,
		})
	}

	return &ShareResults, nil
}
func (r *shareRepositoryGorm) Insert(userOwner uuid.UUID, payload ShareInsertDB) (*uuid.UUID, error) {
	db := r.db
	newTag := entities.ShareEntity{
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		ExpiredDate: payload.ExpiredDate,
		UserShareId: userOwner,
	}
	if err := db.Create(&newTag).Error; err != nil {
		return nil, err
	}
	return &newTag.ShareId, nil

}
func (r *shareRepositoryGorm) GetById(shareId *uuid.UUID) (*[]ShareResultQuery, error) {
	var results []entities.ShareEntity
	db := r.db

	if err := db.Select("share_id, start_date, end_date,expired_date,user_share_id").Find(&results).Error; err != nil {
		return nil, err
	}
	var ShareResults []ShareResultQuery
	for _, result := range results {
		ShareResults = append(ShareResults, ShareResultQuery{
			ShareId:     result.ShareId,
			StartDate:   result.StartDate,
			EndDate:     result.EndDate,
			ExpiredDate: result.ExpiredDate,
			UserShareId: result.UserShareId,
		})
	}

	return &ShareResults, nil

}
