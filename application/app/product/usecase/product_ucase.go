package usecase

import (
	"context"
	"fmt"
	"go-fiber-clean-architecture/application/domain"
)

type ProductUseCase struct {
	MysqlProductRepo domain.ProductRepository
	ElastichProductRepo domain.ProductRepository
}

func NewProductUseCase(mysqlProductRepo domain.ProductRepository, elastichProductRepo domain.ProductRepository) domain.ProductUseCase {
	return &ProductUseCase{
		MysqlProductRepo: mysqlProductRepo,
		ElastichProductRepo: elastichProductRepo,
	}
}

func (uc ProductUseCase) GetByID(ctx context.Context, id int64) (domain.Product, error) {
	// mengambil data dari repository
	res, err := uc.ElastichProductRepo.GetByID(ctx, id)
	// ini adalah contoh penggunaan error handling
	// tapi ini sebenarnya tidak perlu karena error handling sudah di handle di layer delivery
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc ProductUseCase) GetAll(ctx context.Context) ([]domain.Product, error) {
	res, err := uc.ElastichProductRepo.GetAll(ctx)
	return res, err
}

func (uc ProductUseCase) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	product, err := uc.MysqlProductRepo.Create(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}
	res, err := uc.ElastichProductRepo.Create(ctx, product)

	return res, err
}

func (uc ProductUseCase) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	// pengecekan apakah data ada atau tidak
	_, err := uc.ElastichProductRepo.GetByID(ctx, product.ID)
	if err != nil {
		return domain.Product{}, err
	}

	//	pengupdatean data
	product, err = uc.MysqlProductRepo.Update(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}

	res, err := uc.ElastichProductRepo.Update(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}

	fmt.Println(res)
	return product, err
}

func (uc ProductUseCase) Delete(ctx context.Context, id int64) error {
	// pengecekan apakah data ada atau tidak
	_, err := uc.MysqlProductRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// penghapusan data
	err = uc.MysqlProductRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	err = uc.ElastichProductRepo.Delete(ctx, id)
	return err
}