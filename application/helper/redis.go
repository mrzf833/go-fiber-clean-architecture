package helper

import (
	"context"
	"flag"
	"github.com/gofiber/storage/redis/v3"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/helper/helper2"
	"strconv"
)

func ConnectRedis()  {
	// jika MODE test dijalankan
	if flag.Lookup("test.v") != nil {
		config.RedisPort = helper2.GetEnv("REDIS_PORT_TEST", "6379")
		config.RedisHost = helper2.GetEnv("REDIS_HOST_TEST", "localhost")
		config.RedisUser = helper2.GetEnv("REDIS_USER_TEST", "")
		config.RedisPassword = helper2.GetEnv("REDIS_PASSWORD_TEST", "")
		config.RedisDbName = helper2.GetEnv("REDIS_DB_TEST", "1")
	}

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
