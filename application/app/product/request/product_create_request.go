package request


type ProductCreateRequest struct {
	Name string `json:"name" validate:"required"`
	CategoryId int64 `json:"category_id" validate:"required"`
}