package repositories

import (
	"money-service/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type incomeRepositoryGorm struct {
	db *gorm.DB
}

func NewGormIncomeRepository(db *gorm.DB) IncomeRepository {
	return &incomeRepositoryGorm{db: db}
}
func (r *incomeRepositoryGorm) CreateIncome(id uuid.UUID, payload IncomeInsertDb) (*uuid.UUID, error) {

	return nil, nil
}

func (r *incomeRepositoryGorm) GetIncomesByUser(userId uuid.UUID) (*[]IncomeResultQuery, error) {
	var results []entities.IncomesEntity
	db := r.db

	if err := db.Select("").Where("user_onwer = ?", userId).Find(&results).Error; err != nil {
		return nil, err
	}
	var incomesResults []IncomeResultQuery
	for _, result := range results {
		incomesResults = append(incomesResults, IncomeResultQuery{
			IncomeId: result.IncomeId,
		})
	}

	return &incomesResults, nil
}
