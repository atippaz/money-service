package controllers

import (
	repositories "money-service/src/repositories/user_share"
	user_share "money-service/src/services/user_share"
	jwtUtils "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func NewFiberUserShareController(service user_share.UserShareService) FiberUserShareController {
	return &UsershareController{service}
}
func (s UsershareController) GetAllUserShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetAll(&claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []UserShareResult{}
		for _, result := range *res {
			results = append(results, UserShareResult{
				UserShareId: result.UserShareId,
				UserId:      result.UserId,
				ShareId:     result.ShareId,
			})
		}
		return c.JSON(results)

	}
}
func (s UsershareController) InsertUserShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := UserShareRequest{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.Insert(claims.UserId, repositories.UserShareInsertDB{
			UserId: payload.UserShareId,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(res)
	}
}
