package repositories

import (
	"money-service/src/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type expenseRepositoryGorm struct {
	db *gorm.DB
}

func NewGormExpenseRepository(db *gorm.DB) ExpenseRepository {
	return &expenseRepositoryGorm{db: db}
}

func (r *expenseRepositoryGorm) GetExpensesByUser(userId uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]ExpenseResultQuery, error) {
	var results []entities.ExpensesEntity
	db := r.db
	query := db.Where("user_owner = ?", userId)

	if startDate != nil {
		query = query.Where("created_date >= ?", *startDate)
	}

	if endDate != nil {
		query = query.Where("created_date <= ?", *endDate)
	}
	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}
	var incomesResults []ExpenseResultQuery
	for _, result := range results {
		incomesResults = append(incomesResults, ExpenseResultQuery{
			ExpenseId:   result.ExpenseId,
			CreatedDate: result.CreatedDate,
			TagId:       result.TagId,
			Value:       result.Value,
			UserOwner:   result.UserOwner,
		})
	}

	return &incomesResults, nil
}

func (r *expenseRepositoryGorm) CreateExpense(userId uuid.UUID, payload ExpenseInsertDb) (*uuid.UUID, error) {
	db := r.db
	newExpense := entities.ExpensesEntity{
		ExpenseId:   uuid.New(),
		TagId:       payload.TagId,
		Value:       payload.Value,
		UserOwner:   userId,
		CreatedDate: payload.Date,
	}
	if err := db.Create(&newExpense).Error; err != nil {
		return nil, err
	}
	return &newExpense.ExpenseId, nil
}
