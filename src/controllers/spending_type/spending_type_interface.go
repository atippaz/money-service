package controllers

import (
	spendingType "money-service/src/services/spending_type"

	"github.com/gofiber/fiber/v2"
)

type SpendingTypeController[T any] interface {
	GetSpendingHandler() T
}

type spendingTypeController struct {
	service spendingType.SpendingTypeService
}

type FiberSpendingTypeController interface {
	SpendingTypeController[fiber.Handler]
}

type SpendingTypeResult struct {
	NameEn         string `json:"nameEn" `
	NameTh         string `json:"nameTh" `
	SpendingTypeId string `json:"spendingTypeId" `
}
