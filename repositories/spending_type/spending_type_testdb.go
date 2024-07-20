package repositories

import (
	"money-service/config"
	"money-service/interfaces"
)

type spendingRepositoryTest struct {
	db config.DummySpendingInstance
}

func GormTestSpendingTypeRepository(db config.DummySpendingInstance) SpendingTypeRepository {
	return &spendingRepositoryTest{db: db}
}

func (r *spendingRepositoryTest) GetSpendingTypes() (*[]interfaces.SpendingTypeResultQuery, error) {

	db := r.db
	results := db.GetSpending()
	var spendingTypeResults []interfaces.SpendingTypeResultQuery
	for _, result := range results {
		spendingTypeResults = append(spendingTypeResults, interfaces.SpendingTypeResultQuery{
			SpendingTypeId: result.ID,
			NameTh:         result.Name,
			NameEn:         result.Name,
		})
	}

	return &spendingTypeResults, nil
}
