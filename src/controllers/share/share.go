package controllers

import (
	Share "money-service/src/services/share"

	"github.com/gofiber/fiber/v2"
)

func NewFiberShareController(service Share.ShareService) FiberShareController {
	return &shareController{service}
}
func (s shareController) GetAllShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// res, err := s.service.GetShares()
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// results := []ShareResult{}
		// for _, result := range *res {
		// 	results = append(results, ShareResult{
		// 		NameTh:  result.NameTh,
		// 		NameEn:  result.NameEn,
		// 		ShareId: result.ShareId.String(),
		// 	})
		// }
		return c.JSON(nil)

	}
}
func (s shareController) InsertShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// res, err := s.service.GetShares()
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// results := []ShareResult{}
		// for _, result := range *res {
		// 	results = append(results, ShareResult{
		// 		NameTh:  result.NameTh,
		// 		NameEn:  result.NameEn,
		// 		ShareId: result.ShareId.String(),
		// 	})
		// }
		return c.JSON(nil)
	}
}
func (s shareController) GetByIdShareHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// res, err := s.service.GetShares()
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// results := []ShareResult{}
		// for _, result := range *res {
		// 	results = append(results, ShareResult{
		// 		NameTh:  result.NameTh,
		// 		NameEn:  result.NameEn,
		// 		ShareId: result.ShareId.String(),
		// 	})
		// }
		return c.JSON(nil)

	}
}
