package repositories

import (
	"fmt"
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type expenseRepositoryGorm struct {
	db *gorm.DB
}

func ExpenseRepositoryGorm(db *gorm.DB) IExpenseRepository {
	return &expenseRepositoryGorm{db: db}
}

func (r *expenseRepositoryGorm) GetUserById(id string) (*interfaces.UserResult, error) {
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
