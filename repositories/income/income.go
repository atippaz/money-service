package repositories

import (
	"money-service/interfaces"

	"github.com/google/uuid"
)

type IIncomeRepository interface {
	GetIncomesByUser(userId uuid.UUID) (*[]interfaces.IncomeResultQuery, error)
	CreateIncome(id uuid.UUID, payload interfaces.IncomeInsertDb) (*uuid.UUID, error)
}
