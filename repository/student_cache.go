package repository

import (
	"belajar-redis/model/web"
	"context"
	"github.com/go-redis/redis/v8"
)

type StudentCache interface {
	SetAll(ctx context.Context, client *redis.Client, response []web.StudentResponse)
	FindAll(ctx context.Context, client *redis.Client) ([]web.StudentResponse, error)
	SetById(ctx context.Context, client *redis.Client, response web.StudentResponse)
	FindById(ctx context.Context, client *redis.Client, studentId uint) (web.StudentResponse, error)
	Delete(ctx context.Context, client *redis.Client, studentId uint)
}
