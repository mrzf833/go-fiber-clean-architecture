package usecase

import (
	"context"
	"encoding/csv"
	"go-fiber-clean-architecture/application/config"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/utils"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

type categoryUseCase struct {
	CategoryRepo domain.CategoryRepository
}

func NewCategoryUseCase(categoryRepo domain.CategoryRepository) domain.CategoryUseCase {
	return &categoryUseCase{
		CategoryRepo: categoryRepo,
	}
}

func (uc *categoryUseCase) GetByID(ctx context.Context, id int64) (domain.Category, error) {
	// mengambil data dari repository
	res, err := uc.CategoryRepo.GetByID(ctx, id)
	// ini adalah contoh penggunaan error handling
	// tapi ini sebenarnya tidak perlu karena error handling sudah di handle di layer delivery
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *categoryUseCase) GetAll(ctx context.Context) ([]domain.Category, error) {
	res, err := uc.CategoryRepo.GetAll(ctx)
	return res, err
}

func (uc *categoryUseCase) Create(ctx context.Context, category domain.Category) (domain.Category, error) {
	res, err := uc.CategoryRepo.Create(ctx, category)
	return res, err
}

func (uc *categoryUseCase) Update(ctx context.Context, category domain.Category) (domain.Category, error) {
	// pengecekan apakah data ada atau tidak
	_, err := uc.CategoryRepo.GetByID(ctx, category.ID)
	if err != nil {
		return category, err
	}

	//	pengupdatean data
	res, err := uc.CategoryRepo.Update(ctx, category)
	return res, err
}

func (uc *categoryUseCase) Delete(ctx context.Context, id int64) error {
	// pengecekan apakah data ada atau tidak
	_, err := uc.CategoryRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// penghapusan data
	err = uc.CategoryRepo.Delete(ctx, id)
	return err
}

//func (uc *categoryUseCase) CreateWithCsv(ctx context.Context, file io.Reader) error {
//	reader := csv.NewReader(file)
//	reader.FieldsPerRecord = -1 // Allow variable number of fields
//	data, err := reader.ReadAll()
//	if err != nil {
//		return err
//	}
//
//	// skip data field csv
//	data = data[1:]
//
//	totalRecords := len(data)
//
//	batchGoRoutineSize := 4000
//
//	for i := 0; i < totalRecords; i += batchGoRoutineSize {
//		end := i + batchGoRoutineSize
//
//		if end > totalRecords {
//			end = totalRecords
//		}
//
//		batch := data[i:end]
//		// do something
//		go uc.createWithBatch(ctx, batch, 1000)
//	}
//
//	return err
//}

func (uc *categoryUseCase) CreateWithCsv(ctx context.Context, file io.Reader, idTrackerCategory int64) {
	defer utils.Recover()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// skip data field csv
	data = data[1:]

	totalRecords := len(data)

	batchGoRoutineSize := 4000

	for i := 0; i < totalRecords; i += batchGoRoutineSize {
		end := i + batchGoRoutineSize

		if end > totalRecords {
			end = totalRecords
		}

		batch := data[i:end]
		// do something
		go uc.createWithBatch(ctx, batch, 1000, idTrackerCategory)
	}
}

func (uc *categoryUseCase)createWithBatch(ctx context.Context, data [][]string, size int, idTrackerCategory int64)  {
	defer utils.Recover()
	var categoryRecords []domain.Category
	for _, record := range data {
		rows := strings.Split(record[0], ";")
		dataPush := domain.Category{
			Name: rows[0],
		}
		categoryRecords = append(categoryRecords, dataPush)
	}

	err := uc.CategoryRepo.CreateInBatches(ctx, categoryRecords, size)
	if err != nil {
		log.Fatal(err)
	}

	config.DB.Updates(&domain.TrackerCategory{
		ID: idTrackerCategory,
		Name: "Category Berubah",
		Now: "",
		End: strconv.FormatInt(time.Now().UnixMilli(), 10),
	})
}