package request

type FileSaveCreateRequest struct {
	File []byte `validate:"required,file"`
}
