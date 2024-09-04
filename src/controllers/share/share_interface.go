package controllers

import (
	share "money-service/src/services/share"

	"github.com/gofiber/fiber/v2"
)

type ShareController[T any] interface {
	GetAllShareHandler() T
	GetByIdShareHandler() T
	InsertShareHandler() T
}

type shareController struct {
	service share.ShareService
}

type FiberShareController interface {
	ShareController[fiber.Handler]
}

type ShareResult struct {
	NameEn      string `json:"nameEn" `
	NameTh      string `json:"nameTh" `
	ShareTypeId string `json:"ShareTypeId" `
}
