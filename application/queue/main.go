package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"log"
)

func main()  {
	example()
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379", DB: 2},
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
	//mux.HandleFunc(ExampleTask, HandleTaskExample)
	//mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())
	mux.HandleFunc(TypeEmailDelivery, HandleEmailDeliveryTask)
	// ...register other handlers...

	fmt.Println(srv)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	// Email delivery code ...
	return nil
}

func example()  {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
		DB: 2,
	})

	defer client.Close()

	task, err := NewEmailDeliveryTask(42, "welcome_email")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

const (
	ExampleTask   = "message:deliver"
	TypeEmailDelivery   = "email:deliver"
)