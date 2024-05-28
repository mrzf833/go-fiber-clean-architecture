package app

import (
	"flag"
	"github.com/hibiken/asynq"
	job_redis "go-fiber-clean-architecture/application/app/category/job/redis"
	gorm_mysql "go-fiber-clean-architecture/application/app/category/repository/gorm/mysql"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/utils"
	"go-fiber-clean-architecture/application/utils/helper2"
	"log"
	"strconv"
)

func AppRun()  {
	utils.SetLogger()

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

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: config.QueueHost + ":" + config.QueuePort,
			DB: queueDb,
			Username: config.QueueUsername,
			Password: config.QueuePassword,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()

	// connect db
	utils.ConnectDB()

	// setup category repository for queue category
	repositoryCatgeory := gorm_mysql.NewMysqlCategoryRepository(config.DB)
	queueCategory := job_redis.NewQueueCategory(repositoryCatgeory)
	mux.HandleFunc(job_redis.TypeCategoryWithCsvQueue, queueCategory.HandleCategoryCreateWithCsvQueue)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
