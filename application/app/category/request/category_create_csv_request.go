package request

import "mime/multipart"

type CategoryCreateCsvRequest struct {
	File *multipart.FileHeader `form:"file" validate:"required,file_custom_validate"`
}
