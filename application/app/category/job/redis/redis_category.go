package job_redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/utils"
	"log"
	"strings"
)

type CategoryWithCsvQueue struct {
	Data [][]string
	Size int
}

const TypeCategoryWithCsvQueue = "category_with_csv_queue"

func NewCategoryCreateWithCsvQueue(data [][]string, size int) (*asynq.Task, error) {
	payload, err := json.Marshal(CategoryWithCsvQueue{
		Data: data,
		Size: size,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeCategoryWithCsvQueue, payload), nil
}

type QueueCategory struct {
	repo domain.CategoryRepository
}

func NewQueueCategory(repo domain.CategoryRepository) *QueueCategory {
	return &QueueCategory{repo: repo}
}

func (q *QueueCategory) HandleCategoryCreateWithCsvQueue(ctx context.Context, t *asynq.Task) error {
	var payload CategoryWithCsvQueue
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed on HandleCategoryCreateWithCsvQueue : %v: %w", err, asynq.SkipRetry)
	}

	defer utils.Recover()
	var categoryRecords []domain.Category
	for _, record := range payload.Data {
		rows := strings.Split(record[0], ";")
		dataPush := domain.Category{
			Name: rows[0],
		}
		categoryRecords = append(categoryRecords, dataPush)
	}

	err := q.repo.CreateInBatches(ctx, categoryRecords, payload.Size)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success create category with csv queue")
	return nil
}