package repositories

import (
	"money-service/entities"
	"money-service/interfaces"

	"gorm.io/gorm"
)

type incomeRepositoryGorm struct {
	db *gorm.DB
}

func IncomeRepositoryGorm(db *gorm.DB) IIncomeRepository {
	return &incomeRepositoryGorm{db: db}
}

func (r *incomeRepositoryGorm) GetIncomesByUser(userId string) (*[]interfaces.IncomeResultQuery, error) {
	var results []entities.IncomesEntity
	db := r.db

	if err := db.Select("spending_type_id, name_th, name_en").First(&results).Error; err != nil {
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
