package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/config"
)

func JwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.JwtKey)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	})
}

//func AuthMiddleware() fiber.Handler {
//	return JwtMiddleware()
//}

func AuthMiddleware() []fiber.Handler {
	return []fiber.Handler{
		JwtMiddleware(),CheckTokenOnRedis,
	}
}

func CheckTokenOnRedis(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	// get token from redis
	dataTokenByte, err := config.RedisDb.Get(username.(string))
	if err != nil {
		return err
	}

	if len(dataTokenByte) > 0 {
		return c.Next()
	}
	return c.Status(401).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}