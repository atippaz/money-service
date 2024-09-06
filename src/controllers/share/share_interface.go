package controllers

import (
	share "money-service/src/services/share"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	ShareId     uuid.UUID `json:"shareId" `
	StartDate   time.Time `json:"startDate" `
	EndDate     time.Time `json:"endDate" `
	ExpiredDate time.Time `json:"expiredDate" `
	UserShareId uuid.UUID `json:"userShareId" `
}

type ShareRequest struct {
	StartDate   time.Time `json:"startDate" `
	EndDate     time.Time `json:"endDate" `
	ExpiredDate time.Time `json:"expiredDate" `
}
