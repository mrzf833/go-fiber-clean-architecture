package mocks

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/mock"
)

type CategoryWithCsvQueue struct {
	mock.Mock
}

type QueueCategory struct {
	mock.Mock
}

func NewCategoryCreateWithCsvQueue(data [][]string, size int) (*asynq.Task, error) {
	mockStruct := CategoryWithCsvQueue{}
	args := mockStruct.Called(data, size)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*asynq.Task), args.Error(1)
}

func (q *QueueCategory) HandleCategoryCreateWithCsvQueue(ctx context.Context, t *asynq.Task) error {
	args := q.Called(ctx, t)
	return args.Error(0)
}
