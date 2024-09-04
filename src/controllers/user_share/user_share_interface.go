package controllers

import (
	user_share "money-service/src/services/user_share"

	"github.com/gofiber/fiber/v2"
)

type UserShareController[T any] interface {
	GetAllUserShareHandler() T
	InsertUserShareHandler() T
}

type UsershareController struct {
	service user_share.UserShareService
}

type FiberUserShareController interface {
	UserShareController[fiber.Handler]
}

type UserShareResult struct {
	NameEn          string `json:"nameEn" `
	NameTh          string `json:"nameTh" `
	UserShareTypeId string `json:"UserShareTypeId" `
}
