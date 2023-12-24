package usecase

import (
	"context"
	"go-fiber-clean-architecture/application/domain"
)

type categoryUseCase struct {
	categoryRepo domain.CategoryRepository
}

func NewCategoryUseCase(categoryRepo domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *categoryUseCase) GetByID(ctx context.Context, id int64) (domain.Category, error) {
	// mengambil data dari repository
	res, err := uc.categoryRepo.GetByID(ctx, id)
	// ini adalah contoh penggunaan error handling
	// tapi ini sebenarnya tidak perlu karena error handling sudah di handle di layer delivery
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *categoryUseCase) GetAll(ctx context.Context) ([]domain.Category, error) {
	res, err := uc.categoryRepo.GetAll(ctx)
	return res, err
}

func (uc *categoryUseCase) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	res, err := uc.categoryRepo.Create(ctx, category)
	return res, err
}

func (uc *categoryUseCase) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	// pengecekan apakah data ada atau tidak
	_, err := uc.categoryRepo.GetByID(ctx, category.ID)
	if err != nil {
		return category, err
	}

	//	pengupdatean data
	res, err := uc.categoryRepo.Update(ctx, category)
	return res, err
}

func (uc *categoryUseCase) Delete(ctx context.Context, id int64) error {
	// pengecekan apakah data ada atau tidak
	_, err := uc.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// penghapusan data
	err = uc.categoryRepo.Delete(ctx, id)
	return err
}