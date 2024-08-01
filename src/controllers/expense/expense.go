package controllers

import (
	"fmt"
	expenseSevice "money-service/src/services/expense"
	Datetime "money-service/src/utils/datetime"
	jwt "money-service/src/utils/jwt"
	jwtUtils "money-service/src/utils/jwt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewFiberExpenseController(service expenseSevice.ExpenseService, datetime Datetime.Datetime) FiberExpenseController {
	return &expenseController{service, datetime}
}

func (s expenseController) GetExpensesByUser() fiber.Handler {
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
		res, err := s.service.GetExpensesByUser(claims.UserId, startDate, endDate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []ExpenseResult{}
		for _, result := range *res {
			results = append(results, ExpenseResult{
				CreatedDate: result.CreatedDate,
				ExpenseId:   result.ExpenseId.String(),
				TagId:       result.TagId.String(),
				UserOwner:   result.UserOwner.String(),
				Value:       result.Value,
			})
		}
		return c.JSON(results)
	}
}

func (s expenseController) CreateExpense() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var payload ExpenseInsertRequest
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}

		claims := c.Locals("user").(*jwt.AuthClaims)
		res, err := s.service.CreateExpense(claims.UserId, expenseSevice.ExpenseInsert{
			TagId: payload.TagId,
			Value: payload.Value,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.JSON(res)
	}
}
func (s expenseController) GetSummary() fiber.Handler {

	return func(c *fiber.Ctx) error {
		monthStr := c.Query("month")
		yearStr := c.Query("year")
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			fmt.Println("Error parsing month:", err)
			c.Status(fiber.ErrBadRequest.Code).SendString("invalid Month")
		}
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			fmt.Println("Error parsing year:", err)
			c.Status(fiber.ErrBadRequest.Code).SendString("invalid Year")
		}
		startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

		endOfMonth := time.Date(year, time.Month(month)+1, 0, 23, 59, 59, 999999999, time.UTC)

		claims := c.Locals("user").(*jwtUtils.AuthClaims)
		res, err := s.service.GetSummary(claims.UserId, &startOfMonth, &endOfMonth)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		results := []ExpenseSummaryResult{}
		for _, result := range *res {
			results = append(results, ExpenseSummaryResult{
				TagId: result.TagId.String(),
				Value: result.Value,
			})
		}
		return c.JSON(results)
	}
}
