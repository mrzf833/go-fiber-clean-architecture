package usecase

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/helper"
)

type fileSaveUseCase struct {
	fileSaveRepo domain.FileSaveRepository
}

func NewFileSaveUseCase(fileSaveRepo domain.FileSaveRepository) domain.FileSaveUsecase {
	return &fileSaveUseCase{
		fileSaveRepo: fileSaveRepo,
	}
}

func (uc *fileSaveUseCase) GetAll(c *fiber.Ctx) ([]domain.FileSave, error) {
	res, err := uc.fileSaveRepo.GetAll(c)
	return res, err
}

func (uc *fileSaveUseCase) Create(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	// insert data ke database menggunakan gorm
	file, err := c.FormFile("file")
	if err != nil {
		return domain.FileSave{}, err
	}

	// Save file to root directory:
	applicationPath := helper.GetApplicationPath()
	err = helper.SaveFile(c, file, applicationPath+"/storage/public/upload_file", file.Filename)
	if err != nil {
		return domain.FileSave{}, err
	}

	fileSave.Name = "/upload_file/" + file.Filename
	res, err := uc.fileSaveRepo.Create(c, fileSave)
	return res, err
}

func (uc *fileSaveUseCase) Update(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	//	pengupdatean data
	res, err := uc.fileSaveRepo.Update(c, fileSave)
	return res, err
}

func (uc *fileSaveUseCase) Delete(ctx *fiber.Ctx, id int64) error {
	// penghapusan data
	err := uc.fileSaveRepo.Delete(ctx, id)
	return err
}