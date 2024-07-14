package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type systemTagRepositoryGorm struct {
	db *gorm.DB
}

func SystemTagRepositoryGorm(db *gorm.DB) ISystemTagRepository {
	return &systemTagRepositoryGorm{db: db}
}

func (r *systemTagRepositoryGorm) GetAllSystemTags() (*[]interfaces.SystemTagResultQuery, error) {
	var results []entities.SystemTagEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	var systemTagResults []interfaces.SystemTagResultQuery
	for _, result := range results {
		systemTagResults = append(systemTagResults, interfaces.SystemTagResultQuery{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			TagId:          result.TagId,
			IsActive:       result.IsActive,
		})
	}

	return &systemTagResults, nil
}
