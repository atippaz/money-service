package repositories

import "money-service/interfaces"

type IExpenseRepository interface {
	GetExpensesByUser(userId string) (*[]interfaces.ExpenseResultQuery, error)
}
