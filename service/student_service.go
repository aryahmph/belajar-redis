package service

import (
	"belajar-redis/model/web"
	"context"
)

type StudentService interface {
	Create(ctx context.Context, request web.StudentCreateRequest) (web.StudentResponse, error)
	Update(ctx context.Context, request web.StudentUpdateRequest) (web.StudentResponse, error)
	Delete(ctx context.Context, studentNim string) error
	FindByNim(ctx context.Context, studentNim string) (web.StudentResponse, error)
	FindAll(ctx context.Context) []web.StudentResponse
}
