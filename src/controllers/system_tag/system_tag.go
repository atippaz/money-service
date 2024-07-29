package controllers

import (
	SystemTagService "money-service/src/services/system_tag"

	"github.com/gofiber/fiber/v2"
)

func NewFiberSystemTagController(service SystemTagService.SystemTagService) FiberSystemTagController {
	return &systemTagController{service}
}
func (s systemTagController) GetAllSystemTags() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := s.service.GetAllSystemTags()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []TagsResult{}
		for _, result := range *res {
			results = append(results, TagsResult{
				TagId:          result.TagId.String(),
				NameTh:         result.NameTh,
				NameEn:         result.NameEn,
				IsActive:       result.IsActive,
				SpendingTypeId: result.SpendingTypeId.String(),
			})
		}
		return c.JSON(results)
	}
}
