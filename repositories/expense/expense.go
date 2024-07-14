package repositories

import (
	"money-service/interfaces"

	"github.com/google/uuid"
)

type IExpenseRepository interface {
	GetExpensesByUser(userId string) (*[]interfaces.ExpenseResultQuery, error)
	CreateExpense(userId string, payload interfaces.ExpenseResultInsertDb) (*uuid.UUID, error)
}
