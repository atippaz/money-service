package controllers

import (
	authSevice "money-service/services/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthController[T any] interface {
	Register() T
	Login() T
	Logout() T
}
type authController struct {
	service authSevice.AuthService
}
type FiberAuthController interface {
	AuthController[fiber.Handler]
}
