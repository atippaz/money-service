package initials

import (
	"money-service/config"
	"money-service/controllers"
	"money-service/middlewares"
	custom_tag "money-service/repositories/custom_tag"
	expense "money-service/repositories/expense"
	income "money-service/repositories/income"
	spending_type "money-service/repositories/spending_type"
	system_tag "money-service/repositories/system_tag"
	user "money-service/repositories/user"
	"money-service/utils"

	"money-service/routes"
	"money-service/services"

	"github.com/gofiber/fiber/v2"
)

var spendingRepo spending_type.SpendingTypeRepository
var expenseRepo expense.ExpenseRepository
var incomeRepo income.IncomeRepository
var system_tagRepo system_tag.SystemTagRepository
var custom_tagRepo custom_tag.CustomTagRepository
var userRepo user.UserRepository

var spendingService *services.SpendingTypeService
var expenseService *services.ExpenseService
var incomeService *services.IncomeService
var system_tagService *services.SystemTagService
var custom_tagService *services.CustomTagService
var userService *services.UserService
var authService *services.AuthService

var spendingController controllers.FiberSpendingTypeController
var expenseController controllers.FiberExpenseController
var incomeController controllers.FiberIncomeController
var systemTagController controllers.FiberSystemTagController
var customTagController controllers.FiberCustomTagController
var userController controllers.FiberUserController
var authController controllers.FiberAuthController

var encodeUtils utils.Hasher
var jwtUtils utils.Jwt

var middleWareJwt middlewares.JWTMiddleware

func initUtils(config *config.Config) {
	encodeUtils = utils.NewHasher()
	jwtUtils = utils.NewJwt(config.JWTSECERT)
}
func initServices() {
	spendingService = services.NewSpendingTypeService(spendingRepo)
	expenseService = services.NewExpenseService(expenseRepo)
	incomeService = services.NewIncomeService(incomeRepo)
	system_tagService = services.NewSystemTagService(system_tagRepo)
	custom_tagService = services.NewCustomTagService(custom_tagRepo)
	userService = services.NewUserService(userRepo, encodeUtils)
	authService = services.NewAuthService(userService, encodeUtils, jwtUtils)
}

func initContollers() {
	spendingController = controllers.NewFiberSpendingTypeController(spendingService)
	expenseController = controllers.NewFiberExpenseController(expenseService)
	incomeController = controllers.NewFiberIncomeController(incomeService)
	systemTagController = controllers.NewFiberSystemTagController(system_tagService)
	customTagController = controllers.NewFiberCustomTagController(custom_tagService)
	userController = controllers.NewFiberUserController(userService)
	authController = controllers.NewFiberAuthController(authService)
}

func initMiddleWare(config *config.Config) {
	middleWareJwt = middlewares.NewJwtMiddleware(jwtUtils)
}
func initRoutes(app fiber.Router) {
	routes.SpendingTypeRoute(app, spendingController)
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
	initMiddleWare(config)
	initRoutes(app)
}
