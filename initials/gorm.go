package initials

import (
	custom_tag "money-service/repositories/custom_tag"
	expense "money-service/repositories/expense"
	income "money-service/repositories/income"
	spendingType "money-service/repositories/spending_type"
	system_tag "money-service/repositories/system_tag"
	user "money-service/repositories/user"

	"gorm.io/gorm"
)

func InitialGorm(db_gorm *gorm.DB) {
	spendingRepo = spendingType.SpendingTypeRepositoryGorm(db_gorm)
	userRepo = user.UserRepositoryGorm(db_gorm)
	expenseRepo = expense.ExpenseRepositoryGorm(db_gorm)
	incomeRepo = income.IncomeRepositoryGorm(db_gorm)
	system_tagRepo = system_tag.SystemTagRepositoryGorm(db_gorm)
	custom_tagRepo = custom_tag.CustomTagRepositoryGorm(db_gorm)
}
