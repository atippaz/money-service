package middlewares

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func FormatDataMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			return err
		}

		statusCode := c.Response().StatusCode()
		body := c.Response().Body()

		response := Response{
			StatusCode: statusCode,
			Message:    "success",
			Data:       nil,
		}

		if statusCode >= 400 {
			var errorResponse map[string]interface{}
			if err := json.Unmarshal(body, &errorResponse); err == nil {
				if msg, ok := errorResponse["error"].(string); ok {
					response.Message = msg
				} else {
					response.Message = "An error occurred"
				}
			} else {
				response.Message = "An error occurred"
			}
		} else {
			var temp interface{}
			fmt.Println(body)
			if err := json.Unmarshal(body, &temp); err == nil {
				response.Data = temp
			}
			response.Message = "success"
		}

		modifiedBody, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshalling modified response body:", err)
			return nil
		}

		c.Response().SetBody(modifiedBody)

		return nil
	}
}
