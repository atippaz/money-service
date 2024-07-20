package repositories

import (
	"money-service/src/config"
)

type spendingRepositoryTest struct {
	db config.DummySpendingInstance
}

func GormTestSpendingTypeRepository(db config.DummySpendingInstance) SpendingTypeRepository {
	return &spendingRepositoryTest{db: db}
}

func (r *spendingRepositoryTest) GetSpendingTypes() (*[]SpendingTypeResultQuery, error) {

	db := r.db
	results := db.GetSpending()
	var spendingTypeResults []SpendingTypeResultQuery
	for _, result := range results {
		spendingTypeResults = append(spendingTypeResults, SpendingTypeResultQuery{
			SpendingTypeId: result.ID,
			NameTh:         result.Name,
			NameEn:         result.Name,
		})
	}

	return &spendingTypeResults, nil
}
