package service

import (
	"belajar-redis/model/web"
	"context"
)

type StudentService interface {
	Create(ctx context.Context, request web.StudentCreateRequest) (web.StudentResponse, error)
	Update(ctx context.Context, request web.StudentUpdateRequest) (web.StudentResponse, error)
	Delete(ctx context.Context, studentNim string)
	FindByNim(ctx context.Context, categoryId int) web.StudentResponse
	FindAll(ctx context.Context) []web.StudentResponse
}
