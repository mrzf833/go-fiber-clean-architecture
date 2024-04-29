package helper

import (
	"context"
	"github.com/gofiber/storage/redis/v3"
	"go-fiber-clean-architecture/application/config"
	"strconv"
)

func ConnectRedis()  {
	// convert string to int port
	port,err := strconv.Atoi(config.RedisPort)
	if err != nil {
		panic("failed to convert port redis to int")
	}
	// convert string to int select db
	dbRedis,err := strconv.Atoi(config.RedisDbName)
	if err != nil {
		panic("failed to convert db redis to int ")
	}
	config.RedisDb = redis.New(redis.Config{
		Host:     config.RedisHost,
		Port:     port,
		Username: config.RedisUser,
		Password: config.RedisPassword,
		Database: dbRedis,
	})

	// check connection to redis
	if pong := config.RedisDb.Conn().Ping(context.Background()); pong.String() != "ping: PONG" {
		panic("failed to connect redis")
	}
}
