package controllers

import (
	"fmt"
	incomeService "money-service/src/services/income"
	Datetime "money-service/src/utils/datetime"
	jwtUtils "money-service/src/utils/jwt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewFiberIncomeController(service incomeService.IncomeService, datetime Datetime.Datetime) FiberIncomeController {
	return &incomeController{service, datetime}
}
func (s incomeController) CreateIncome() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// //todo get uuid from user middleware
		// //todo get payload form body
		// payload := IncomeInsertDb{}
		// id := uuid.New()
		// res, err := s.service.CreateIncome(id, payload)
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		// }
		// return c.JSON(res)
		return c.JSON("")
	}
}
func (s incomeController) GetIncomesByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var startDate *time.Time
		var endDate *time.Time
		startDateStr := c.Query("startDate")
		endDateStr := c.Query("endDate")
		if startDateStr != "" {
			date, err := s.datetime.ConvertToDate(startDateStr)
			if err != nil {
				date, _ = s.datetime.GetStartDate()
			}
			startDate = &date
		}
		if endDateStr != "" {
			date, err := s.datetime.ConvertToDate(endDateStr)
			fmt.Println(err)
			if err != nil {
				date, _ = s.datetime.GetEndDate()
			}
			endDate = &date
		}
		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		fmt.Println(endDate)
		res, err := s.service.GetIncomesByUser(claims.UserId, startDate, endDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		results := []IncomeResult{}
		for _, result := range *res {
			results = append(results, IncomeResult{
				CreatedDate: result.CreatedDate,
				IncomeId:    result.IncomeId.String(),
				TagId:       result.TagId.String(),
				UserOwner:   result.UserOwner.String(),
				Value:       result.Value,
			})
		}
		return c.JSON(results)
	}
}
