package initials

import (
	"money-service/controllers"
	custom_tag "money-service/repositories/custom_tag"
	expense "money-service/repositories/expense"
	income "money-service/repositories/income"
	spending_type "money-service/repositories/spending_type"
	system_tag "money-service/repositories/system_tag"
	user "money-service/repositories/user"

	"money-service/routes"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

var spendingRepo spending_type.ISpendingTypeRepository
var expenseRepo expense.IExpenseRepository
var incomeRepo income.IIncomeRepository
var system_tagRepo system_tag.ISystemTagRepository
var custom_tagRepo custom_tag.ICustomTagRepository
var userRepo user.IUserRepository

var spendingService *services.ISpendingTypeService
var expenseService *services.IExpenseService
var incomeService *services.IIncomeService
var system_tagService *services.ISystemTagService
var custom_tagService *services.ICustomTagService
var userService *services.IUserService

var spendingController controllers.ISpendingTypeController
var expenseController controllers.IExpenseController
var incomeController controllers.IIncomeController
var systemTagController controllers.ISystemTagController
var customTagController controllers.ICustomTagController
var userController controllers.IUserController

func initServices() {
	spendingService = services.SpendingTypeService(spendingRepo)
	expenseService = services.ExpenseService(expenseRepo)
	incomeService = services.IncomeService(incomeRepo)
	system_tagService = services.SystemTagService(system_tagRepo)
	custom_tagService = services.CustomTagService(custom_tagRepo)
	userService = services.UserService(userRepo)
}

func initContollers() {
	spendingController = controllers.SpendingTypeController(spendingService)
	expenseController = controllers.ExpenseController(expenseService)
	incomeController = controllers.IncomeController(incomeService)
	systemTagController = controllers.SystemTagController(system_tagService)
	customTagController = controllers.CustomTagController(custom_tagService)
	userController = controllers.UserController(userService)
}

func initRoutes(app fiber.Router) {
	routes.SpendingTypeRoute(app, spendingController)
	routes.ExpenseRoute(app, expenseController)
	routes.IncomeRoute(app, incomeController)
	routes.SystemTagRoute(app, systemTagController)
	routes.CustomTagRoute(app, customTagController)
	routes.UserRoute(app, userController)
}

func InitApplication(app fiber.Router) {
	initServices()
	initContollers()
	initRoutes(app)
}
