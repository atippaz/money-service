package controllers

import (
	repositories "money-service/src/repositories/share"
	Share "money-service/src/services/share"
	jwtUtils "money-service/src/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

func NewFiberShareController(service Share.ShareService) FiberShareController {
	return &shareController{service}
}
func (s shareController) GetAllShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)

		res, err := s.service.GetAll(&claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []ShareResult{}
		for _, result := range *res {
			results = append(results, ShareResult{
				ShareId:     result.ShareId,
				StartDate:   result.StartDate,
				EndDate:     result.EndDate,
				ExpiredDate: result.ExpiredDate,
				UserShareId: result.ShareId,
			})
		}
		return c.JSON(results)

	}
}
func (s shareController) InsertShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := ShareRequest{}
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.Insert(claims.UserId, repositories.ShareInsertDB{
			StartDate:   payload.StartDate,
			EndDate:     payload.EndDate,
			ExpiredDate: payload.ExpiredDate,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s shareController) GetByIdShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetById(&claims.UserId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []ShareResult{}
		for _, result := range *res {
			results = append(results, ShareResult{
				ShareId:     result.ShareId,
				StartDate:   result.StartDate,
				EndDate:     result.EndDate,
				ExpiredDate: result.ExpiredDate,
				UserShareId: result.ShareId,
			})
		}
		return c.JSON(results)

	}
}
