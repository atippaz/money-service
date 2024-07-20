package repositories

import (
	"money-service/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type customTagRepositoryGorm struct {
	db *gorm.DB
}

func NewGormCustomTagRepository(db *gorm.DB) CustomTagRepository {
	return &customTagRepositoryGorm{db: db}
}
func (r *customTagRepositoryGorm) CreateCustomTag(userOwner uuid.UUID, payload CustomTagInsertDB) (*uuid.UUID, error) {
	db := r.db
	_payload := entities.CustomTagEntity{
		NameTh:         payload.NameTh,
		NameEn:         payload.NameEn,
		IsActive:       true,
		SpendingTypeId: payload.SpendingTypeId,
		UserOwner:      userOwner,
	}
	if err := db.Create(&_payload).Error; err != nil {
		return nil, err
	}
	return &_payload.TagId, nil
}

func (r *customTagRepositoryGorm) GetCustomTagsByUser(ownerId string) (*[]CustomTagResultQuery, error) {
	var results []entities.CustomTagEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	var customTagResults []CustomTagResultQuery
	for _, result := range results {
		customTagResults = append(customTagResults, CustomTagResultQuery{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			TagId:          result.TagId,
			IsActive:       result.IsActive,
			UserOwner:      result.UserOwner,
		})
	}

	return &customTagResults, nil
}
