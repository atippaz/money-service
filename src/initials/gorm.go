package initials

import (
	custom_tag "money-service/src/repositories/custom_tag"
	expense "money-service/src/repositories/expense"
	income "money-service/src/repositories/income"
	spendingType "money-service/src/repositories/spending_type"
	system_tag "money-service/src/repositories/system_tag"
	user "money-service/src/repositories/user"

	"gorm.io/gorm"
)

func InitialGorm(db_gorm *gorm.DB) {
	spendingRepo = spendingType.NewGormSpendingTypeRepository(db_gorm)
	userRepo = user.NewGormUserRepository(db_gorm)
	expenseRepo = expense.NewGormExpenseRepository(db_gorm)
	incomeRepo = income.NewGormIncomeRepository(db_gorm)
	systemTagRepo = system_tag.NewGormSystemTagRepository(db_gorm)
	customTagRepo = custom_tag.NewGormCustomTagRepository(db_gorm)
}
