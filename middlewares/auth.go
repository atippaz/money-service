package middlewares

import (
	"fmt"
	"money-service/config"

	"github.com/gofiber/fiber/v2"
)

// Custom claims struct
type Claims struct {
	Username string `json:"userName"`
	Email    string `json:"email"`
	UserId   string `json:"userId"`
	// jwt.RegisteredClaims
}

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

		// claims := &Claims{}
		// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (Claims{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		// 	}
		// 	return []byte("your-secret-key"), nil
		// })

		// if err != nil || !token.Valid {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": "Invalid or expired JWT",
		// 	})
		// }

		c.Locals("user", config.UserId)

		return c.Next()
	}
}
