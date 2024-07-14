package repositories

import "money-service/interfaces"

type IExpenseRepository interface {
	GetUserById(id string) (*interfaces.UserResult, error)
}
