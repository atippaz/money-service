package repositories

import (
	"fmt"
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type spendingRepositoryGorm struct {
	db *gorm.DB
}

func NewGormSpendingTypeRepository(db *gorm.DB) SpendingTypeRepository {
	return &spendingRepositoryGorm{db: db}
}

func (r *spendingRepositoryGorm) GetSpendingTypes() (*[]interfaces.SpendingTypeResultQuery, error) {
	var results []entities.SpendingTypeEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Print(results)
	var spendingTypeResults []interfaces.SpendingTypeResultQuery
	for _, result := range results {
		spendingTypeResults = append(spendingTypeResults, interfaces.SpendingTypeResultQuery{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
		})
	}

	return &spendingTypeResults, nil
}
