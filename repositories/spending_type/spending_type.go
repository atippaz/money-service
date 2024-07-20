package repositories

import "money-service/interfaces"

type SpendingTypeRepository interface {
	GetSpendingTypes() (*[]interfaces.SpendingTypeResultQuery, error)
}
