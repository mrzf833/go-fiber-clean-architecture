package usecase

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-clean-architecture/application/app/auth/request"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/exception"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type authUseCase struct {
	authRepo domain.AuthRepository
	userRepo domain.UserRepository
}

func NewAuthUseCase(authRepo domain.AuthRepository, userRepo domain.UserRepository) domain.AuthUseCase {
	return &authUseCase{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}

func (uc *authUseCase) Login(c *fiber.Ctx, request request.AuthCreateRequest) (map[string]interface{}, error) {
	// mengambil data dari repository
	user, err := uc.userRepo.GetByUsername(c.Context(), request.Username)
	// ini adalah contoh penggunaan error handling
	// tapi ini sebenarnya tidak perlu karena error handling sudah di handle di layer delivery
	if err != nil {
		return map[string]interface{}{}, exception.NewHandlerCustomError(400, fiber.Map{
			"message": "username or password is wrong",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	//pengecekan username dan password apakah sama atau tidak
	if user.Username != request.Username || err != nil{
		return map[string]interface{}{}, exception.NewHandlerCustomError(400, fiber.Map{
			"message": "username or password is wrong",
		})
	}
	// data auth ke penyimpanan
	dataAuth := domain.Auth{
		Username: user.Username,
	}

	duration := config.ExpireToken

	// generate token
	token, err := uc.generateToken(map[string]any{
		"username": user.Username,
	}, duration)
	if err != nil {
		return map[string]interface{}{}, err
	}

	// set token ke dataAuth
	dataAuth.Token = token
	// store data ke redis
	data, err := uc.authRepo.CreateToken(context.TODO(), dataAuth, duration)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return map[string]interface{}{
		"token": data.Token,
		"expire": data.Expire,
		"username": user.Username,
	}, nil
}

func (uc *authUseCase) generateToken(data map[string]any, duration time.Duration) (string, error) {
	// set expire token
	day := time.Now().Add(duration).Unix()
	// membuat token
	claims := jwt.MapClaims{
		"username": data["username"],
		"exp": day,
	}

	// generate encoded token and send it as response.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", exception.ErrInternalServerError
	}

	// set data auth token ke struct
	return t, nil
}

func (uc *authUseCase) User(c *fiber.Ctx) jwt.MapClaims{
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}

func (uc *authUseCase) Logout(c *fiber.Ctx) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	// delete token di redis
	uc.authRepo.DeleteToken(context.TODO(), username.(string))
}