package app

import (
	"flag"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/utils/helper2"
	"log"
	"net/http"
	"strconv"
)

func AppRun()  {
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

	h := asynqmon.New(asynqmon.Options{
		RootPath: "/", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{
			Addr: config.QueueHost + ":" + config.QueuePort,
			DB: queueDb,
			Username: config.QueueUsername,
			Password: config.QueuePassword,
		},
	})

	// Note: We need the tailing slash when using net/http.ServeMux.
	http.Handle(h.RootPath()+"/", h)

	// Go to http://localhost:8080/monitoring to see asynqmon homepage.
	host := helper2.GetEnv("QUEUE_MONITORING_HOST", "localhost")
	port := helper2.GetEnv("QUEUE_MONITORING_PORT", "3000")

	fmt.Println("Server is running on " + host + ":" + port)
	err = http.ListenAndServe(host + ":" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
