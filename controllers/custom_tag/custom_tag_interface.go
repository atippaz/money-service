package controllers

import (
	customTagSevice "money-service/services/custom_tag"

	"github.com/gofiber/fiber/v2"
)

type CustomTagController[T any] interface {
	GetCustomTagsByUser() T
	CreateCustomTag() T
}
type customTagController struct {
	service customTagSevice.CustomTagService
}

type FiberCustomTagController interface {
	CustomTagController[fiber.Handler]
}
