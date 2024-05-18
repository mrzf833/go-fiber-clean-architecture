package config

import (
	"github.com/gofiber/storage/redis/v3"
	"go-fiber-clean-architecture/application/utils/helper2"
)

var (
	RedisDb *redis.Storage
	RedisPort = helper2.GetEnv("REDIS_PORT", "6379")
	RedisHost = helper2.GetEnv("REDIS_HOST", "localhost")
	RedisUser = helper2.GetEnv("REDIS_USER", "")
	RedisPassword = helper2.GetEnv("REDIS_PASSWORD", "")
	RedisDbName = helper2.GetEnv("REDIS_DB", "0")
)