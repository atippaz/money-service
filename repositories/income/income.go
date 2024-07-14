package repositories

import "money-service/interfaces"

type IIncomeRepository interface {
	GetIncomesByUser(userId string) (*[]interfaces.IncomeResultQuery, error)
}
