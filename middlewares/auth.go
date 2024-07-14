package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
)

// Custom claims struct
type Claims struct {
	Username string `json:"username"`
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
		// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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

		// c.Locals("user", claims.Username)

		return c.Next()
	}
}
