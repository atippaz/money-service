package controllers

import (
	SystemTagService "money-service/src/services/system_tag"

	"github.com/gofiber/fiber/v2"
)

type SystemTagController[T any] interface {
	GetAllSystemTags() T
}
type systemTagController struct {
	service SystemTagService.SystemTagService
}

type FiberSystemTagController interface {
	SystemTagController[fiber.Handler]
}

type TagsResult struct {
	IsActive       bool   `json:"isActive" `
	NameEn         string `json:"nameEn" `
	NameTh         string `json:"nameTh" `
	SpendingTypeId string `json:"spendingTypeId" `
	TagId          string `json:"tagId" `
}
