package utils

import (
	"flag"
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
	})

	defer client.Close()

	return client
}

func SetClientQueue() {
	if config.CientQueue == nil {
		config.CientQueue = NewClientQueueRedis()
	}
}
