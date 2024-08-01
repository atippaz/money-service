package repositories

import (
	"money-service/src/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type incomeRepositoryGorm struct {
	db *gorm.DB
}

func NewGormIncomeRepository(db *gorm.DB) IncomeRepository {
	return &incomeRepositoryGorm{db: db}
}
func (r *incomeRepositoryGorm) CreateIncome(userId uuid.UUID, payload IncomeInsertDb) (*uuid.UUID, error) {
	db := r.db
	newIncome := entities.IncomesEntity{
		IncomeId:    uuid.New(),
		TagId:       payload.TagId,
		Value:       payload.Value,
		UserOwner:   userId,
		CreatedDate: payload.Date,
	}
	if err := db.Create(&newIncome).Error; err != nil {
		return nil, err
	}
	return (&newIncome.IncomeId), nil
}

func (r *incomeRepositoryGorm) GetIncomesByUser(userId uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeResultQuery, error) {
	var results []entities.IncomesEntity
	db := r.db
	query := db.Where("user_owner = ?", userId)

	if startDate != nil {
		query = query.Where("created_date >= ?", *startDate)
	}

	if endDate != nil {
		query = query.Where("created_date <= ?", *endDate)
	}
	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}
	var incomesResults []IncomeResultQuery
	for _, result := range results {
		incomesResults = append(incomesResults, IncomeResultQuery{
			IncomeId:    result.IncomeId,
			CreatedDate: result.CreatedDate,
			TagId:       result.TagId,
			Value:       result.Value,
			UserOwner:   result.UserOwner,
		})
	}

	return &incomesResults, nil
}

func (r *incomeRepositoryGorm) GetSummary(userId uuid.UUID, startDate *time.Time, endDate *time.Time) (*[]IncomeResultQuery, error) {
	var results []entities.IncomesEntity
	r.GetIncomesByUser(userId, startDate, endDate)
	db := r.db
	query := db.Where("user_owner = ?", userId)

	if startDate != nil {
		query = query.Where("created_date >= ?", *startDate)
	}

	if endDate != nil {
		query = query.Where("created_date <= ?", *endDate)
	}
	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}
	var incomesResults []IncomeResultQuery
	for _, result := range results {
		incomesResults = append(incomesResults, IncomeResultQuery{
			IncomeId:    result.IncomeId,
			CreatedDate: result.CreatedDate,
			TagId:       result.TagId,
			Value:       result.Value,
			UserOwner:   result.UserOwner,
		})
	}

	return &incomesResults, nil
}
