package controllers

import (
	userService "money-service/services/user"

	"github.com/gofiber/fiber/v2"
)

type UserController[T any] interface {
	GetUserById() T
	DeActiveAccount() T
}
type userController struct {
	service userService.UserService
}
type FiberUserController interface {
	UserController[fiber.Handler]
}
