package repositories

import (
	"fmt"
	"money-service/entities"

	"gorm.io/gorm"
)

type systemTagRepositoryGorm struct {
	db *gorm.DB
}

func NewGormSystemTagRepository(db *gorm.DB) SystemTagRepository {
	return &systemTagRepositoryGorm{db: db}
}

func (r *systemTagRepositoryGorm) GetAllSystemTags() (*[]SystemTagResultQuery, error) {
	var results []entities.SystemTagEntity
	db := r.db

	if err := db.Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Println(results)
	var systemTagResults []SystemTagResultQuery
	for _, result := range results {
		systemTagResults = append(systemTagResults, SystemTagResultQuery{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
			TagId:          result.TagId,
			IsActive:       result.IsActive,
		})
	}

	return &systemTagResults, nil
}
