package web

type StudentResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type StudentCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

type StudentUpdateRequest struct {
	Id   uint   `validate:"required"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}
