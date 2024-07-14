package repositories

import (
	"fmt"
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type incomeRepositoryGorm struct {
	db *gorm.DB
}

func IncomeRepositoryGorm(db *gorm.DB) IIncomeRepository {
	return &incomeRepositoryGorm{db: db}
}

func (r *incomeRepositoryGorm) GetUserById(id string) (*interfaces.UserResult, error) {
	var result entities.UserEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").First(&result).Error; err != nil {
		return nil, err
	}
	fmt.Print(result)
	var spendingTypeResults = interfaces.UserResult{
		UserId: result.UserId,
	}

	return &spendingTypeResults, nil
}
