package request

import "mime/multipart"

type FileSaveCreateRequest struct {
	File *multipart.FileHeader `form:"file" validate:"required,file_custom_validate"`
	Coba string `form:"coba" validate:"required"`
}
