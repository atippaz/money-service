package repositories

import (
	"money-service/src/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type expenseRepositoryGorm struct {
	db *gorm.DB
}

func NewGormExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepositoryGorm{db: db}
}

func (r *expenseRepositoryGorm) GetExpensesByUser(userId uuid.UUID) (*[]ExpenseResultQuery, error) {
	var results []entities.ExpensesEntity
	db := r.db

	if err := db.Find(&results).Error; err != nil {
		return nil, err
	}
	expenseResults := make([]ExpenseResultQuery, 0)
	for _, result := range results {
		expenseResults = append(expenseResults, ExpenseResultQuery{
			ExpenseId: result.ExpenseId,
		})
	}

	return &expenseResults, nil
}
func (r *expenseRepositoryGorm) CreateExpense(userId uuid.UUID, payload ExpenseInsertDb) (*uuid.UUID, error) {
	db := r.db
	newExpense := entities.ExpensesEntity{
		ExpenseId: uuid.New(),
		TagId:     payload.TagId,
		Value:     payload.Value,
		UserOwner: userId,
	}
	if err := db.Create(&newExpense).Error; err != nil {
		return nil, err
	}
	return &newExpense.ExpenseId, nil
}
