package job_redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/utils"
	"strings"
)

type CategoryWithCsvQueue struct {
	Data [][]string
	Size int
}

const TypeCategoryWithCsvQueue = "category_with_csv_queue"

func NewCategoryCreateWithCsvQueue(data [][]string, size int) (*asynq.Task, error) {
	payload, err := json.Marshal(&CategoryWithCsvQueue{
		Data: data,
		Size: size,
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return asynq.NewTask(TypeCategoryWithCsvQueue, payload), nil
}

type QueueCategory struct {
	Repo domain.CategoryRepository
}

func NewQueueCategory(repo domain.CategoryRepository) domain.QueueCategoryJob {
	return &QueueCategory{Repo: repo}
}

func (q *QueueCategory) HandleCategoryCreateWithCsvQueue(ctx context.Context, t *asynq.Task) error {
	var payload CategoryWithCsvQueue
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		log.Errorf("json.Unmarshal failed on HandleCategoryCreateWithCsvQueue : %v: %v", err, asynq.SkipRetry)
		return err
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

	err := q.Repo.CreateInBatches(ctx, categoryRecords, payload.Size)
	if err != nil {
		log.Error(err)
		return err
	}

	fmt.Println("Success create category with csv queue")
	return nil
}