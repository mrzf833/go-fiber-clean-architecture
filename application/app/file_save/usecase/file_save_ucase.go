package usecase

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-clean-architecture/application/domain"
	"go-fiber-clean-architecture/application/helper"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
	"time"
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

func (uc *fileSaveUseCase) GetByID(ctx *fiber.Ctx, id int64) (domain.FileSave, error) {
	// mengambil data dari repository
	res, err := uc.fileSaveRepo.GetByID(ctx, id)
	// ini adalah contoh penggunaan error handling
	// tapi ini sebenarnya tidak perlu karena error handling sudah di handle di layer delivery
	if err != nil {
		return res, err
	}
	return res, nil
}

func (uc *fileSaveUseCase) Create(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	// insert data ke database menggunakan gorm
	file, err := c.FormFile("file")
	if err != nil {
		return domain.FileSave{}, err
	}

	// Save file to root directory:
	applicationPath := helper.GetApplicationPath()
	nameUnix := strconv.FormatInt(time.Now().UnixMicro(), 10)
	fileName := nameUnix + filepath.Ext(file.Filename)
	err = helper.SaveFile(c, file, applicationPath+"/storage/public/upload_file", fileName)
	if err != nil {
		return domain.FileSave{}, err
	}
	fileSave.Name = "/upload_file/" + fileName
	res, err := uc.fileSaveRepo.Create(c, fileSave)
	if err != nil {
		err2 := os.Remove(applicationPath + "/storage/public/" + fileSave.Name)
		if err2 != nil {
			return domain.FileSave{}, err2
		}
	}
	return res, err
}

//func (uc *fileSaveUseCase) Update(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
//	file, err := c.FormFile("file")
//	if err != nil {
//		return domain.FileSave{}, err
//	}
//
//	// get file before update
//	fileSaveBefore, err := uc.fileSaveRepo.GetByID(c, fileSave.ID)
//	if err != nil {
//		return domain.FileSave{}, err
//	}
//
//	res := domain.FileSave{}
//	err = uc.fileSaveRepo.GetDb().Transaction(func(tx *gorm.DB) error {
//
//
//		return nil
//	})
//
//	// Save file to root directory:
//	applicationPath := helper.GetApplicationPath()
//	nameUnix := strconv.FormatInt(time.Now().UnixMicro(), 10)
//	// set file name and path
//	fileName := nameUnix + filepath.Ext(file.Filename)
//	// save file new
//	err = helper.SaveFile(c, file, applicationPath+"/storage/public/upload_file", fileName)
//	if err != nil {
//		return domain.FileSave{}, err
//	}
//
//	// set file name to file save
//	fileSave.Name = "/upload_file/" + fileName
//	// hapus file lama
//	err = os.Remove(applicationPath + "/storage/public/" + fileSaveBefore.Name)
//	// jika gagal menghapus file lama maka akan mengembalikan error dan mengahapus file baru
//	if err != nil {
//		// menghapus file baru
//		err2 := os.Remove(applicationPath + "/storage/public/" + fileSave.Name)
//		if err2 != nil {
//			return domain.FileSave{}, err2
//		}
//		return domain.FileSave{}, err
//	}
//
//	res, err = uc.fileSaveRepo.Update(c, fileSave)
//	if err != nil {
//		return domain.FileSave{}, err
//	}
//	//	pengupdatean data
//	return res, err
//}

func (uc *fileSaveUseCase) Update(c *fiber.Ctx, fileSave domain.FileSave) (domain.FileSave, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return domain.FileSave{}, err
	}

	// get file before update
	fileSaveBefore, err := uc.fileSaveRepo.GetByID(c, fileSave.ID)
	if err != nil {
		return domain.FileSave{}, err
	}

	res := domain.FileSave{}
	err = uc.fileSaveRepo.GetDb().Transaction(func(tx *gorm.DB) error {

		// Save file to root directory:
		applicationPath := helper.GetApplicationPath()
		nameUnix := strconv.FormatInt(time.Now().UnixMicro(), 10)
		// set file name and path
		fileName := nameUnix + filepath.Ext(file.Filename)
		// save file new
		err = helper.SaveFile(c, file, applicationPath+"/storage/public/upload_file", fileName)
		if err != nil {
			return err
		}

		// set file name to file save
		fileSave.Name = "/upload_file/" + fileName
		res, err = uc.fileSaveRepo.Update(c, fileSave)
		if err != nil {
			// jika gagal update maka akan menghapus file baru
			err2 := os.Remove(applicationPath + "/storage/public/" + fileSave.Name)
			if err2 != nil {
				return err2
			}

			return err
		}
		// hapus file lama
		err = os.Remove(applicationPath + "/storage/public/" + fileSaveBefore.Name)
		// jika gagal menghapus file lama maka akan mengembalikan error dan menghapus file baru
		if err != nil {
			// menghapus file baru
			err2 := os.Remove(applicationPath + "/storage/public/" + fileSave.Name)
			if err2 != nil {
				return err2
			}
			return err
		}
		return nil
	})
	//	pengupdatean data
	return res, err
}

func (uc *fileSaveUseCase) Delete(ctx *fiber.Ctx, id int64) error {
	// penghapusan data
	fileBefore, err := uc.fileSaveRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = uc.fileSaveRepo.GetDb().Transaction(func(tx *gorm.DB) error {
		err = uc.fileSaveRepo.Delete(ctx, id)
		if err != nil {
			return err
		}

		applicationPath := helper.GetApplicationPath()
		err = os.Remove(applicationPath + "/storage/public" + fileBefore.Name)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}