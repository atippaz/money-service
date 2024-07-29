package controllers

import (
	userService "money-service/src/services/user"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

type UserInfo struct {
	UserId      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	UserName    string    `json:"userName"`
	LastName    string    `json:"lastName"`
	FirstName   string    `json:"firstName"`
	DisplayName string    `json:"displayName"`
	Profile     string    `json:"profile"`
	CreatedDate time.Time `json:"createdDate"`
}
