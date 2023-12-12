package usecase

import (
	"context"
	"go-fiber-clean-architecture/application/domain"
)

type categoryUseCase struct {
	
}

func NewCategoryUseCase() domain.CategoryUseCase {
	return &categoryUseCase{}
}

func (uc *categoryUseCase) GetByID(ctx context.Context, id int64) (domain.Category, error) {
	return domain.Category{}, nil
}