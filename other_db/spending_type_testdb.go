package other_db

import (
	"fmt"
	"money-service/interfaces"
	"money-service/repositories"
)

type spendingRepositoryTest struct {
	db DummySpendingInstance
}

func SpendingTypeRepositoryTest(db DummySpendingInstance) repositories.ISpendingTypeRepository {
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
