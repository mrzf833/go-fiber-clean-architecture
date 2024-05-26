package config

import (
	"github.com/hibiken/asynq"
	"go-fiber-clean-architecture/application/utils/helper2"
)

var (
	CientQueue *asynq.Client
	QueueHost = helper2.GetEnv("QUEUE_HOST", "localhost")
	QueuePort = helper2.GetEnv("QUEUE_PORT", "6379")
	QueueDb = helper2.GetEnv("QUEUE_DB", "0")
	QueuePassword = helper2.GetEnv("QUEUE_PASSWORD", "")
	QueueUsername = helper2.GetEnv("QUEUE_USERNAME", "")
)