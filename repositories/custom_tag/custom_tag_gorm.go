package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type customTagRepositoryGorm struct {
	db *gorm.DB
}

func CustomTagRepositoryGorm(db *gorm.DB) ICustomTagRepository {
	return &customTagRepositoryGorm{db: db}
}

func (r *customTagRepositoryGorm) GetCustomTagsByUser(ownerId string) (*[]interfaces.CustomTagResultQuery, error) {
	var results []entities.CustomTagEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	var customTagResults []interfaces.CustomTagResultQuery
	for _, result := range results {
		customTagResults = append(customTagResults, interfaces.CustomTagResultQuery{
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
