package controllers

import (
	authSevice "money-service/src/services/auth"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthController[T any] interface {
	Register() T
	Login() T
	Logout() T
}
type authController struct {
	service  authSevice.AuthService
	validate *validator.Validate
}
type FiberAuthController interface {
	AuthController[fiber.Handler]
}
type AuthRegisterRequest struct {
	UserName    string `json:"userName" validate:"required,min=3"`
	Email       string `json:"email" validate:"required,email"`
	LastName    string `json:"lastName" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	DisplayName string `json:"displayName"`
	Password    string `json:"password" validate:"required,min=6"`
	Profile     string `json:"profile"`
}

type AuthLoginRequest struct {
	Credential string `json:"credential" validate:"required,min=3"`
	Password   string `json:"password" validate:"required,min=6"`
}
