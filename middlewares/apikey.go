package middlewares

import (
	"fmt"
	"money-service/config"
	"os"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
)

func ApiKeyMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		fmt.Println("check api key")
		apiKey := os.Getenv("API_KEY")
		fmt.Println("no token ", apiKey, tokenString)
		c.Locals("user", config.UserId)
		return c.Next()
	}
}
