package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type incomeRepositoryGorm struct {
	db *gorm.DB
}

func IncomeRepositoryGorm(db *gorm.DB) IIncomeRepository {
	return &incomeRepositoryGorm{db: db}
}
func (r *incomeRepositoryGorm) CreateIncome(id uuid.UUID, payload interfaces.IncomeInsertDb) (*uuid.UUID, error) {

	return nil, nil
}

func (r *incomeRepositoryGorm) GetIncomesByUser(userId uuid.UUID) (*[]interfaces.IncomeResultQuery, error) {
	var results []entities.IncomesEntity
	db := r.db

	if err := db.Select("").Where("user_onwer = ?", userId).Find(&results).Error; err != nil {
		return nil, err
	}
	var incomesResults []interfaces.IncomeResultQuery
	for _, result := range results {
		incomesResults = append(incomesResults, interfaces.IncomeResultQuery{
			IncomeId: result.IncomeId,
		})
	}

	return &incomesResults, nil
}
