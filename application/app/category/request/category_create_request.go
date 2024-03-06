package request

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

//var CategoryCreateReequest = map[string]any{
//	"name": "required",
//}
