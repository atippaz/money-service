package controllers

import (
	user_share "money-service/src/services/user_share"

	"github.com/gofiber/fiber/v2"
)

func NewFiberUserShareController(service user_share.UserShareService) FiberUserShareController {
	return &UsershareController{service}
}
func (s UsershareController) GetAllUserShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// res, err := s.service.GetUserShares()
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// results := []UserShareResult{}
		// for _, result := range *res {
		// 	results = append(results, UserShareResult{
		// 		NameTh:  result.NameTh,
		// 		NameEn:  result.NameEn,
		// 		UserShareId: result.UserShareId.String(),
		// 	})
		// }
		return c.JSON(nil)

	}
}
func (s UsershareController) InsertUserShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// res, err := s.service.GetUserShares()
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// results := []UserShareResult{}
		// for _, result := range *res {
		// 	results = append(results, UserShareResult{
		// 		NameTh:  result.NameTh,
		// 		NameEn:  result.NameEn,
		// 		UserShareId: result.UserShareId.String(),
		// 	})
		// }
		return c.JSON(nil)
	}
}
