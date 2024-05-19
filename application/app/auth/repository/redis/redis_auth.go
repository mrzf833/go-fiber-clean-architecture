package redis

import (
	"context"
	"encoding/json"
	"github.com/gofiber/storage/redis/v3"
	"go-fiber-clean-architecture/application/domain"
	"time"
)

type redisAuthRepository struct {
	Db *redis.Storage
}

func NewRedisAuthRepository(db *redis.Storage) domain.AuthRepository {
	return &redisAuthRepository{db}
}

func (r *redisAuthRepository) CreateToken(ctx context.Context, auth domain.Auth, exp time.Duration) (domain.Auth, error)  {
	// store data ke redis
	auth.Expire = time.Now().Add(exp)
	dataAuthByte,_ := json.Marshal(auth)
	err := r.Db.Set(auth.Username, dataAuthByte, exp)
	if err != nil {
		return domain.Auth{}, err

	}
	return auth, err
}

func (r *redisAuthRepository) DeleteToken(ctx context.Context, username string) error  {
	err := r.Db.Delete(username)
	return err
}

func (r *redisAuthRepository) GetToken(ctx context.Context, username string) (domain.Auth, error)  {
	// store data ke redis
	dataAuthByte, err := r.Db.Get(username)
	if err != nil {
		return domain.Auth{}, err
	}
	auth := domain.Auth{}
	err = json.Unmarshal(dataAuthByte, &auth)
	if err != nil {
		return domain.Auth{}, err
	}
	return auth, err
}