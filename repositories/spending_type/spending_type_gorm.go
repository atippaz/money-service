package repositories

import (
	"money-service/entities"

	"gorm.io/gorm"
)

type spendingRepositoryGorm struct {
	db *gorm.DB
}

func NewGormSpendingTypeRepository(db *gorm.DB) SpendingTypeRepository {
	return &spendingRepositoryGorm{db: db}
}

func (r *spendingRepositoryGorm) GetSpendingTypes() (*[]SpendingTypeResultQuery, error) {
	var results []entities.SpendingTypeEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	var spendingTypeResults []SpendingTypeResultQuery
	for _, result := range results {
		spendingTypeResults = append(spendingTypeResults, SpendingTypeResultQuery{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
		})
	}

	return &spendingTypeResults, nil
}
