package request


type ProductUpdateRequest struct {
	Name string `json:"name" validate:"required"`
	CategoryId int64 `json:"category_id" validate:"required" mapstructure:"category_id"`
}