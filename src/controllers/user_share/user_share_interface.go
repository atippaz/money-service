package controllers

import (
	user_share "money-service/src/services/user_share"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	UserShareId uuid.UUID `json:"userShareId" `
	UserId      uuid.UUID `json:"userId" `
	ShareId     uuid.UUID `json:"shareId" `
}
type UserShareRequest struct {
	UserShareId uuid.UUID `json:"userShareId"`
}
