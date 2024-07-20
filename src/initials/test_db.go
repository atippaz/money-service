package initials

import (
	"money-service/src/config"
	repositories "money-service/src/repositories/spending_type"
)

func InitialTest(db config.DummySpendingInstance) {
	spendingRepo = repositories.GormTestSpendingTypeRepository(db)
}
