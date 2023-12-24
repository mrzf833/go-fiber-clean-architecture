package request

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
	CobaLagi string `json:"coba_lagi" validate:"required"`
}

//var CategoryCreateReequest = map[string]any{
//	"name": "required",
//}