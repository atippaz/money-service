package initials

import (
	"money-service/config"
	auth_contoller "money-service/controllers/auth"
	custom_tag_contoller "money-service/controllers/custom_tag"
	expense_contoller "money-service/controllers/expense"
	income_contoller "money-service/controllers/income"
	spending_type_contoller "money-service/controllers/spending_type"
	system_tag_contoller "money-service/controllers/system_tag"
	user_contoller "money-service/controllers/user"
	"money-service/middlewares"

	custom_tag_repo "money-service/repositories/custom_tag"
	expense_repo "money-service/repositories/expense"
	income_repo "money-service/repositories/income"
	spending_type_repo "money-service/repositories/spending_type"
	system_tag_repo "money-service/repositories/system_tag"
	user_repo "money-service/repositories/user"

	hasher "money-service/utils/hasher"
	Jwt "money-service/utils/jwt"

	"money-service/routes"

	auth_service "money-service/services/auth"
	custom_tag_service "money-service/services/custom_tag"
	expense_service "money-service/services/expense"
	income_service "money-service/services/income"
	spending_type_service "money-service/services/spending_type"
	system_tag_service "money-service/services/system_tag"
	user_service "money-service/services/user"

	"github.com/gofiber/fiber/v2"
)

var spendingRepo spending_type_repo.SpendingTypeRepository
var expenseRepo expense_repo.ExpenseRepository
var incomeRepo income_repo.IncomeRepository
var systemTagRepo system_tag_repo.SystemTagRepository
var customTagRepo custom_tag_repo.CustomTagRepository
var userRepo user_repo.UserRepository

var spendingService spending_type_service.SpendingTypeService
var expenseService expense_service.ExpenseService
var incomeService income_service.IncomeService
var system_tagService system_tag_service.SystemTagService
var custom_tagService custom_tag_service.CustomTagService
var userService user_service.UserService
var authService auth_service.AuthService

var customTagController custom_tag_contoller.FiberCustomTagController
var expenseController expense_contoller.FiberExpenseController
var incomeController income_contoller.FiberIncomeController
var spendingTypeController spending_type_contoller.FiberSpendingTypeController
var systemTagController system_tag_contoller.FiberSystemTagController
var userController user_contoller.FiberUserController
var authController auth_contoller.FiberAuthController

var encodeUtils hasher.Hasher
var jwtUtils Jwt.Jwt

var middleWareJwt middlewares.JWTMiddleware

func initUtils(config *config.Config) {
	encodeUtils = hasher.NewHasher()
	jwtUtils = Jwt.NewJwt(config.JWTSECERT)
}
func initServices() {
	spendingService = spending_type_service.NewSpendingTypeService(spendingRepo)
	expenseService = expense_service.NewExpenseService(expenseRepo)
	incomeService = income_service.NewIncomeService(incomeRepo)
	system_tagService = system_tag_service.NewSystemTagService(systemTagRepo)
	custom_tagService = custom_tag_service.NewCustomTagService(customTagRepo)
	userService = user_service.NewUserService(userRepo, encodeUtils)
	authService = auth_service.NewAuthService(userService, encodeUtils, jwtUtils)
}

func initContollers() {
	userController = user_contoller.NewFiberUserController(userService)
	expenseController = expense_contoller.NewFiberExpenseController(expenseService)
	incomeController = income_contoller.NewFiberIncomeController(incomeService)
	systemTagController = system_tag_contoller.NewFiberSystemTagController(system_tagService)
	customTagController = custom_tag_contoller.NewFiberCustomTagController(custom_tagService)
	spendingTypeController = spending_type_contoller.NewFiberSpendingTypeController(spendingService)
	authController = auth_contoller.NewFiberAuthController(authService)
}

func initMiddleWare() {
	middleWareJwt = middlewares.NewJwtMiddleware(jwtUtils)
}
func initRoutes(app fiber.Router) {
	routes.SpendingTypeRoute(app, spendingTypeController)
	routes.ExpenseRoute(app, expenseController, middleWareJwt)
	routes.IncomeRoute(app, incomeController)
	routes.SystemTagRoute(app, systemTagController)
	routes.CustomTagRoute(app, customTagController, middleWareJwt)
	routes.UserRoute(app, userController)
	routes.AuthRoute(app, authController)

}

func InitApplication(app fiber.Router, config *config.Config) {
	initUtils(config)
	initServices()
	initContollers()
	initMiddleWare()
	app.Use(middlewares.ApiKeyMiddleware(config.API_KEY))
	initRoutes(app)
}
