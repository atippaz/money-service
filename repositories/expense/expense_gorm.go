package repositories

import (
	"fmt"
	"money-service/entities"
	"money-service/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type expenseRepositoryGorm struct {
	db *gorm.DB
}

func NewGormExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepositoryGorm{db: db}
}

func (r *expenseRepositoryGorm) GetExpensesByUser(userId uuid.UUID) (*[]interfaces.ExpenseResultQuery, error) {
	var results []entities.ExpensesEntity
	db := r.db

	if err := db.Select("*").Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Println(results)
	var expenseResults []interfaces.ExpenseResultQuery
	for _, result := range results {
		expenseResults = append(expenseResults, interfaces.ExpenseResultQuery{
			ExpenseId: result.ExpenseId,
		})
	}

	return &expenseResults, nil
}
func (r *expenseRepositoryGorm) CreateExpense(userId uuid.UUID, payload interfaces.ExpenseInsertDb) (*uuid.UUID, error) {

	return nil, nil
}
