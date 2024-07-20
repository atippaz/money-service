package middlewares

import (
	"fmt"
	Jwt "money-service/utils/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type JWTMiddleware interface {
	MiddleWare() fiber.Handler
}
type jwtMiddleware struct {
	jwt Jwt.Jwt
}

func NewJwtMiddleware(jwt Jwt.Jwt) JWTMiddleware {
	return &jwtMiddleware{jwt: jwt}
}
func (s jwtMiddleware) MiddleWare() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		}
		tokenString := parts[1]

		fmt.Println(tokenString)

		user, err := s.jwt.DecodeJWT(tokenString)
		fmt.Println(err)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		}
		claim := Jwt.AuthClaims{
			Username: user.Username,
			Email:    user.Email,
			UserId:   user.UserId,
		}
		fmt.Println("sss")
		fmt.Println(claim)

		c.Locals("user", &claim)
		return c.Next()
	}
}
