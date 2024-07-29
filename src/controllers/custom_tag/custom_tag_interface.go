package controllers

import (
	customTagSevice "money-service/src/services/custom_tag"

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

type CustomTagsResult struct {
	IsActive       bool   `json:"isActive" `
	NameEn         string `json:"nameEn" `
	NameTh         string `json:"nameTh" `
	SpendingTypeId string `json:"spendingTypeId" `
	TagId          string `json:"tagId" `
	UserOwner      string `json:"userOwner" `
}
