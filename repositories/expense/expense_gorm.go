package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type expenseRepositoryGorm struct {
	db *gorm.DB
}

func ExpenseRepositoryGorm(db *gorm.DB) IExpenseRepository {
	return &expenseRepositoryGorm{db: db}
}

func (r *expenseRepositoryGorm) GetExpensesByUser(userId string) (*[]interfaces.ExpenseResultQuery, error) {
	var results []entities.ExpensesEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").First(&results).Error; err != nil {
		return nil, err
	}
	var expenseResults []interfaces.ExpenseResultQuery
	for _, result := range results {
		expenseResults = append(expenseResults, interfaces.ExpenseResultQuery{
			ExpenseId: result.ExpenseId,
		})
	}

	return &expenseResults, nil
}
func (r *expenseRepositoryGorm) CreateExpense(userId string, payload interfaces.ExpenseResultInsertDb) (*uuid.UUID, error) {

	return nil, nil
}
