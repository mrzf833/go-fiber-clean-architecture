package utils

import (
	"context"
	"flag"
	"github.com/gofiber/storage/redis/v3"
	"github.com/hibiken/asynq"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/utils/helper2"
	"strconv"
)

func NewClientQueueRedis() *asynq.Client {
	if flag.Lookup("test.v") != nil {
		config.QueueDb     = helper2.GetEnv("QUEUE_DB_TEST", "2")
		config.QueueHost   = helper2.GetEnv("QUEUE_HOST_TEST", "localhost")
		config.QueuePort   = helper2.GetEnv("QUEUE_PORT_TEST", "6379")
		config.QueueUsername = helper2.GetEnv("QUEUE_USERNAME_TEST", "")
		config.QueuePassword = helper2.GetEnv("QUEUE_PASSWORD_TEST", "")
	}

	// convert string to int select db
	queueDb,err := strconv.Atoi(config.QueueDb)
	if err != nil {
		panic("failed to convert db queue redis to int ")
	}

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: config.QueueHost + ":" + config.QueuePort,
		DB: queueDb,
		Username: config.QueueUsername,
		Password: config.QueuePassword,
	})

	return client
}

func NewConnectClientQueueRedis() *redis.Storage {
	// jika MODE test dijalankan
	if flag.Lookup("test.v") != nil {
		config.QueueDb     = helper2.GetEnv("QUEUE_DB_TEST", "2")
		config.QueueHost   = helper2.GetEnv("QUEUE_HOST_TEST", "localhost")
		config.QueuePort   = helper2.GetEnv("QUEUE_PORT_TEST", "6379")
		config.QueueUsername = helper2.GetEnv("QUEUE_USERNAME_TEST", "")
		config.QueuePassword = helper2.GetEnv("QUEUE_PASSWORD_TEST", "")
	}

	// convert string to int port
	port,err := strconv.Atoi(config.QueuePort)
	if err != nil {
		panic("failed to convert port redis to int")
	}
	// convert string to int select db
	dbRedis,err := strconv.Atoi(config.QueueDb)
	if err != nil {
		panic("failed to convert db redis to int ")
	}

	redisDb := redis.New(redis.Config{
		Host:     config.QueueHost,
		Port:     port,
		Username: config.QueueUsername,
		Password: config.QueuePassword,
		Database: dbRedis,
	})

	// check connection to redis
	if pong := redisDb.Conn().Ping(context.Background()); pong.String() != "ping: PONG" {
		panic("failed to connect redis")
	}

	return redisDb
}



func SetClientQueue() {
	if config.CientQueue == nil {
		config.CientQueue = NewClientQueueRedis()
	}
}
