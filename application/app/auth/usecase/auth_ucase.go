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

	// set expire token
	day := time.Now().Add(config.ExpireToken).Unix()
	// membuat token
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp": day,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JwtKey))
	if err != nil {
		return map[string]interface{}{}, exception.ErrInternalServerError
	}
	// data auth ke penyimpanan
	dataAuth := domain.Auth{
		Username: user.Username,
		Token: t,
		Expire: time.Now().Add(config.ExpireToken),
	}


	// store data ke redis
	err = uc.authRepo.CreateToken(context.TODO(), user.Username, dataAuth, config.ExpireToken)
	if err != nil {
		return map[string]interface{}{}, err
	}

	return map[string]interface{}{
		"token": t,
		"expire": day,
		"username": user.Username,
	}, nil
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