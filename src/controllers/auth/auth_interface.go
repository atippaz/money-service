package controllers

import (
	authSevice "money-service/src/services/auth"

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
type AuthRegisterRequest struct {
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"firstName"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password"`
	Profile     string `json:"profile"`
}

type AuthLoginRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}
