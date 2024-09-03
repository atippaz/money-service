package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
)

func ApiKeyMiddleware(keySecert string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		segments := strings.Split(c.Path(), "/")
		lastSegment := segments[len(segments)-1]
		if lastSegment == "webhook" {
			return c.Next()
		}
		apiKeyHeader := c.Get("X-API-KEY")
		if keySecert != apiKeyHeader {
			return c.Status(fiber.StatusUnauthorized).SendString("Missing token")
		}
		return c.Next()
	}
}
