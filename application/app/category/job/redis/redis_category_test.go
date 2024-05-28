package job_redis_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	job_redis "go-fiber-clean-architecture/application/app/category/job/redis"
	"go-fiber-clean-architecture/application/app/category/mocks"
	"testing"
)

func TestNewCategoryCreateWithCsvQueue(t *testing.T) {
	_, err := job_redis.NewCategoryCreateWithCsvQueue([][]string{{"test"}}, 1)
	assert.NoError(t, err)
}

func TestQueueCategory_HandleCategoryCreateWithCsvQueue(t *testing.T) {
	mockCategoryRepo := new(mocks.CategoryRepository)
	mockCategoryRepo.On("CreateInBatches", mock.Anything, mock.AnythingOfType("[]domain.Category"), mock.AnythingOfType("int")).Return(nil).Once()
	q := job_redis.NewQueueCategory(mockCategoryRepo)
	queue, err := job_redis.NewCategoryCreateWithCsvQueue([][]string{{"test"}}, 1)
	assert.NoError(t, err)
	err = q.HandleCategoryCreateWithCsvQueue(context.TODO(), queue)
	assert.NoError(t, err)

	mockCategoryRepo.AssertExpectations(t)
}