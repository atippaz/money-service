package initials

import (
	"money-service/config"
	repositories "money-service/repositories/spending_type"
)

func InitialTest(db config.DummySpendingInstance) {
	spendingRepo = repositories.GormTestSpendingTypeRepository(db)
}
