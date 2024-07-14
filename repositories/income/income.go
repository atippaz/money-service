package repositories

import "money-service/interfaces"

type IIncomeRepository interface {
	GetUserById(id string) (*interfaces.UserResult, error)
}
