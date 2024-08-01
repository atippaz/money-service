package initials

import (
	"money-service/src/config"
	auth_contoller "money-service/src/controllers/auth"
	custom_tag_contoller "money-service/src/controllers/custom_tag"
	expense_contoller "money-service/src/controllers/expense"
	income_contoller "money-service/src/controllers/income"
	spending_type_contoller "money-service/src/controllers/spending_type"
	system_tag_contoller "money-service/src/controllers/system_tag"
	user_contoller "money-service/src/controllers/user"
	"money-service/src/middlewares"

	custom_tag_repo "money-service/src/repositories/custom_tag"
	expense_repo "money-service/src/repositories/expense"
	income_repo "money-service/src/repositories/income"
	spending_type_repo "money-service/src/repositories/spending_type"
	system_tag_repo "money-service/src/repositories/system_tag"
	user_repo "money-service/src/repositories/user"

	Datetime "money-service/src/utils/datetime"
	hasher "money-service/src/utils/hasher"
	Jwt "money-service/src/utils/jwt"

	"money-service/src/routes"

	auth_service "money-service/src/services/auth"
	custom_tag_service "money-service/src/services/custom_tag"
	expense_service "money-service/src/services/expense"
	income_service "money-service/src/services/income"
	spending_type_service "money-service/src/services/spending_type"
	system_tag_service "money-service/src/services/system_tag"
	user_service "money-service/src/services/user"

	"github.com/go-playground/validator"
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
var datetime Datetime.Datetime

var middleWareJwt middlewares.JWTMiddleware
var validate = validator.New()

func initUtils(config *config.Config) {
	encodeUtils = hasher.NewHasher()
	jwtUtils = Jwt.NewJwt(config.JWTSECERT)
	datetime = Datetime.NewDatetime()
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
	expenseController = expense_contoller.NewFiberExpenseController(expenseService, datetime)
	incomeController = income_contoller.NewFiberIncomeController(incomeService, datetime)
	systemTagController = system_tag_contoller.NewFiberSystemTagController(system_tagService)
	customTagController = custom_tag_contoller.NewFiberCustomTagController(custom_tagService)
	spendingTypeController = spending_type_contoller.NewFiberSpendingTypeController(spendingService)
	authController = auth_contoller.NewFiberAuthController(authService, validate)
}

func initMiddleWare() {
	middleWareJwt = middlewares.NewJwtMiddleware(jwtUtils)
}
func initRoutes(app fiber.Router) {
	routes.SpendingTypeRoute(app, spendingTypeController)
	routes.ExpenseRoute(app, expenseController, middleWareJwt)
	routes.IncomeRoute(app, incomeController, middleWareJwt)
	routes.SystemTagRoute(app, systemTagController)
	routes.CustomTagRoute(app, customTagController, middleWareJwt)
	routes.UserRoute(app, userController, middleWareJwt)
	routes.AuthRoute(app, authController)

}

func InitApplication(app fiber.Router, config *config.Config) {
	initUtils(config)
	initServices()
	initContollers()
	initMiddleWare()
	// app.Use(middlewares.FormatDataMiddleware())
	app.Use(middlewares.ApiKeyMiddleware(config.API_KEY))
	initRoutes(app)
}
