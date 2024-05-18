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

func (r *redisAuthRepository) CreateToken(ctx context.Context, username string, auth domain.Auth, exp time.Duration) error  {
	dataAuthByte,_ := json.Marshal(auth)
	err := r.Db.Set(username, dataAuthByte, exp)
	return err
}

func (r *redisAuthRepository) DeleteToken(ctx context.Context, username string) error  {
	err := r.Db.Delete(username)
	return err
}