package controllers

import (
	TagSevice "money-service/src/services/tag"

	"github.com/gofiber/fiber/v2"
)

type TagController[T any] interface {
	GetTagsByUser() T
	CreateTag() T
}
type tagController struct {
	service TagSevice.TagService
}

type FiberTagController interface {
	TagController[fiber.Handler]
}

type TagsResult struct {
	IsActive       bool   `json:"isActive" `
	NameEn         string `json:"nameEn" `
	NameTh         string `json:"nameTh" `
	SpendingTypeId string `json:"spendingTypeId" `
	TagId          string `json:"tagId" `
	Owner          string `json:"owner" `
}
