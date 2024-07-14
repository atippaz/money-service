package gorm_repository

import (
	"fmt"
	"money-service/entities"
	"money-service/interfaces"
	"money-service/repositories"

	"gorm.io/gorm"
)

type spendingRepositoryGorm struct {
	db *gorm.DB
}

func SpendingTypeRepositoryGorm(db *gorm.DB) repositories.ISpendingTypeRepository {
	return &spendingRepositoryGorm{db: db}
}

func (r *spendingRepositoryGorm) GetSpendingTypes() (*[]interfaces.SpendingTypeResult, error) {
	var results []entities.SpendingTypeEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Print(results)
	var spendingTypeResults []interfaces.SpendingTypeResult
	for _, result := range results {
		spendingTypeResults = append(spendingTypeResults, interfaces.SpendingTypeResult{
			SpendingTypeId: result.SpendingTypeId,
			NameTh:         result.NameTh,
			NameEn:         result.NameEn,
		})
	}

	return &spendingTypeResults, nil
}
