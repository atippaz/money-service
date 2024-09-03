package initials

import (
	"money-service/src/config"
	auth_contoller "money-service/src/controllers/auth"
	expense_contoller "money-service/src/controllers/expense"
	income_contoller "money-service/src/controllers/income"
	line_contoller "money-service/src/controllers/line"
	spending_type_contoller "money-service/src/controllers/spending_type"
	tag_contoller "money-service/src/controllers/tag"
	user_contoller "money-service/src/controllers/user"
	"money-service/src/middlewares"

	expense_repo "money-service/src/repositories/expense"
	income_repo "money-service/src/repositories/income"
	spending_type_repo "money-service/src/repositories/spending_type"
	tag_repo "money-service/src/repositories/tag"
	user_repo "money-service/src/repositories/user"

	Datetime "money-service/src/utils/datetime"
	hasher "money-service/src/utils/hasher"
	Jwt "money-service/src/utils/jwt"

	"money-service/src/routes"

	auth_service "money-service/src/services/auth"
	expense_service "money-service/src/services/expense"
	income_service "money-service/src/services/income"
	line_service "money-service/src/services/line"
	spending_type_service "money-service/src/services/spending_type"
	tag_service "money-service/src/services/tag"
	user_service "money-service/src/services/user"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var spendingRepo spending_type_repo.SpendingTypeRepository
var expenseRepo expense_repo.ExpenseRepository
var incomeRepo income_repo.IncomeRepository
var tagRepo tag_repo.TagRepository
var userRepo user_repo.UserRepository

var spendingService spending_type_service.SpendingTypeService
var expenseService expense_service.ExpenseService
var incomeService income_service.IncomeService
var tagService tag_service.TagService
var userService user_service.UserService
var authService auth_service.AuthService
var lineService line_service.LineService

var expenseController expense_contoller.FiberExpenseController
var incomeController income_contoller.FiberIncomeController
var spendingTypeController spending_type_contoller.FiberSpendingTypeController
var tagController tag_contoller.FiberTagController
var userController user_contoller.FiberUserController
var authController auth_contoller.FiberAuthController
var lineController line_contoller.FiberLineTypeController

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
	tagService = tag_service.NewTagService(tagRepo)
	userService = user_service.NewUserService(userRepo, encodeUtils)
	authService = auth_service.NewAuthService(userService, encodeUtils, jwtUtils)
	lineService = line_service.NewLineService(authService, userService, expenseService, tagService)
}

func initContollers() {
	userController = user_contoller.NewFiberUserController(userService)
	expenseController = expense_contoller.NewFiberExpenseController(expenseService, datetime)
	incomeController = income_contoller.NewFiberIncomeController(incomeService, datetime)

	tagController = tag_contoller.NewFiberTagController(tagService)
	spendingTypeController = spending_type_contoller.NewFiberSpendingTypeController(spendingService)
	authController = auth_contoller.NewFiberAuthController(authService, validate)
	lineController = line_contoller.NewFiberSpendingTypeController(lineService)

}

func initMiddleWare() {
	middleWareJwt = middlewares.NewJwtMiddleware(jwtUtils)
}
func initRoutes(app fiber.Router) {
	routes.SpendingTypeRoute(app, spendingTypeController)
	routes.ExpenseRoute(app, expenseController, middleWareJwt)
	routes.IncomeRoute(app, incomeController, middleWareJwt)
	routes.CustomTagRoute(app, tagController, middleWareJwt)
	routes.UserRoute(app, userController, middleWareJwt)
	routes.AuthRoute(app, authController)
	routes.LineRoute(app, lineController)

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
