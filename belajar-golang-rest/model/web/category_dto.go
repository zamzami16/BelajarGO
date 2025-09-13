package web

type CategoryResponse struct {
	Id   int
	Name string `json:"name"`
}

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=1" json:"name"`
}