package initials

import (
	expense "money-service/src/repositories/expense"
	income "money-service/src/repositories/income"
	share "money-service/src/repositories/share"
	user_share "money-service/src/repositories/user_share"

	spendingType "money-service/src/repositories/spending_type"
	tag "money-service/src/repositories/tag"

	user "money-service/src/repositories/user"

	"gorm.io/gorm"
)

func InitialGorm(db_gorm *gorm.DB) {
	spendingRepo = spendingType.NewGormSpendingTypeRepository(db_gorm)
	userRepo = user.NewGormUserRepository(db_gorm)
	expenseRepo = expense.NewGormExpenseRepository(db_gorm)
	incomeRepo = income.NewGormIncomeRepository(db_gorm)
	tagRepo = tag.NewGormTagRepository(db_gorm)
	shareRepo = share.NewGormShareRepository(db_gorm)
	userShareRepo = user_share.NewGormUserShareRepository(db_gorm)

}
