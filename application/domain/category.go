package domain

import "context"

type Category struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

type CategoryUseCase interface {
	GetByID(ctx context.Context, id int64) (Category, error)
}