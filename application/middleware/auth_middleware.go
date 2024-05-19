package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/app/auth/repository/redis"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"time"
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

var authRepository domain.AuthRepository
func AuthMiddleware() []fiber.Handler {
	if authRepository == nil {
		authRepository = redis.NewRedisAuthRepository(config.RedisDb)
	}

	return []fiber.Handler{
		JwtMiddleware(),CheckTokenOnRedis,
	}
}

func CheckTokenOnRedis(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	// get token from redis
	dataToken, _ := authRepository.GetToken(c.Context(), username.(string))

	if dataToken.Token == token.Raw && dataToken.Expire.After(time.Now()) {
		return c.Next()
	}
	return c.Status(401).JSON(fiber.Map{
		"message": "Unauthorized",
	})
}