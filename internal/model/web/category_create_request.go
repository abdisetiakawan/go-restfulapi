package web

type CategoryRequest struct {
	Name string `validate:"required,min=1,max=200"`
}