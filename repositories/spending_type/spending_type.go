package repositories

import "money-service/interfaces"

type ISpendingTypeRepository interface {
	GetSpendingTypes() (*[]interfaces.SpendingTypeResultQuery, error)
}
