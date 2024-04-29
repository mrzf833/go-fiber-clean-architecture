package usecase

import (
	"encoding/json"
	"errors"
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
}

func NewAuthUseCase(authRepo domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{
		authRepo: authRepo,
	}
}

func (uc *authUseCase) Login(c *fiber.Ctx, request request.AuthCreateRequest) (map[string]interface{}, error) {
	// mengambil data dari repository
	user, err := uc.authRepo.GetByUsername(c.Context(), request.Username)
	// ini adalah contoh penggunaan error handling
	// tapi ini sebenarnya tidak perlu karena error handling sudah di handle di layer delivery
	if err != nil {
		return map[string]interface{}{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	//pengecekan username dan password apakah sama atau tidak
	if user.Username != request.Username || err != nil{
		return map[string]interface{}{}, errors.New("username or password is wrong")
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
	// data auth ke redis
	dataAuth := map[string]any{
		"username": user.Username,
		"token": t,
		"expire": time.Now().Add(config.ExpireToken),
	}
	dataAuthByte, _ := json.Marshal(dataAuth)
	// store data ke redis
	config.RedisDb.Set(user.Username, dataAuthByte, config.ExpireToken)

	return map[string]interface{}{
		"token": t,
		"expire": day,
		"username": user.Username,
	}, nil
}
