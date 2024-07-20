package repositories

import (
	"money-service/interfaces"

	"github.com/google/uuid"
)

type ExpenseRepository interface {
	GetExpensesByUser(userId uuid.UUID) (*[]interfaces.ExpenseResultQuery, error)
	CreateExpense(userId uuid.UUID, payload interfaces.ExpenseInsertDb) (*uuid.UUID, error)
}
