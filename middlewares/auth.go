package middlewares

import (
	"fmt"
	"money-service/config"
	"money-service/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// JWTMiddleware function to check token
func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the token from the header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			fmt.Println("no token")
			// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			// 	"message": "Missing or malformed JWT",
			// })
		}
		fmt.Println(tokenString)

		claims := &interfaces.AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}
			return []byte("test"), nil
		})
		fmt.Println(token)

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		}

		c.Locals("user", config.UserId)

		return c.Next()
	}
}
