package controllers

import (
	line "money-service/src/services/line"

	"github.com/gofiber/fiber/v2"
)

type LineController[T any] interface {
	Webhook() T
}

type lineController struct {
	service line.LineService
}

type FiberLineTypeController interface {
	LineController[fiber.Handler]
}
