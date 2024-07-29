package controllers

import (
	customTagSevice "money-service/src/services/custom_tag"
	jwtUtils "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func NewFiberCustomTagController(service customTagSevice.CustomTagService) FiberCustomTagController {
	return &customTagController{service}
}
func (s customTagController) GetCustomTagsByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetCustomTagsByUser(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		var results []CustomTagsResult
		for _, result := range *res {
			results = append(results, CustomTagsResult{
				TagId:          result.TagId.String(),
				NameTh:         result.NameTh,
				NameEn:         result.NameEn,
				IsActive:       result.IsActive,
				SpendingTypeId: result.SpendingTypeId.String(),
				UserOwner:      result.UserOwner.String(),
			})
		}

		return c.JSON(&results)
	}
}
func (s customTagController) CreateCustomTag() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.CreateCustomTag(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
