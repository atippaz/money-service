package repositories

import (
	"fmt"
	"money-service/config"
	"money-service/interfaces"
)

type spendingRepositoryTest struct {
	db config.DummySpendingInstance
}

func SpendingTypeRepositoryTest(db config.DummySpendingInstance) ISpendingTypeRepository {
	return &spendingRepositoryTest{db: db}
}

func (r *spendingRepositoryTest) GetSpendingTypes() (*[]interfaces.SpendingTypeResult, error) {

	db := r.db
	results := db.GetSpending()
	fmt.Print(results)
	var spendingTypeResults []interfaces.SpendingTypeResult
	for _, result := range results {
		spendingTypeResults = append(spendingTypeResults, interfaces.SpendingTypeResult{
			SpendingTypeId: result.ID,
			NameTh:         result.Name,
			NameEn:         result.Name,
		})
	}

	return &spendingTypeResults, nil
}
