package controllers

import (
	"fmt"
	"log"
	line "money-service/src/services/line"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewFiberSpendingTypeController(service line.LineService) FiberLineTypeController {
	return &lineController{service}
}
func (s *lineController) Webhook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		Line := line.LineMessage{}
		fmt.Println(c.Body())

		if err := c.BodyParser(&Line); err != nil {
			log.Println("error parsing body:", err)
			return c.Status(http.StatusOK).SendString("error")
		}

		if len(Line.Events) <= 0 {
			return c.Status(http.StatusOK).SendString("ok")
		}

		go s.service.CommandManager(&Line)

		log.Println("%% message success")
		return c.Status(http.StatusOK).SendString("ok")
	}
}
