package middlewares

import (
	Jwt "money-service/src/utils/jwt"
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
		user, err := s.jwt.DecodeJWT(tokenString)
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

		c.Locals("user", &claim)
		return c.Next()
	}
}
