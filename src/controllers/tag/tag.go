package controllers

import (
	TagSevice "money-service/src/services/tag"
	jwtUtils "money-service/src/utils/jwt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewFiberTagController(service TagSevice.TagService) FiberTagController {
	return &tagController{service}
}
func (s tagController) GetTagsByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		hasSystemStr := c.Query("hasSystem", "false")
		hasSystem, err := strconv.ParseBool(hasSystemStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid hasSystem parameter")
		}
		res, err := s.service.GetTagsByUser(claims.UserId, hasSystem)
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
				Owner:          result.Owner.String(),
			})
		}

		return c.JSON(&results)
	}
}
func (s tagController) CreateTag() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.CreateTag(claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
