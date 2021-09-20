package repository

import (
	"belajar-redis/model/domain"
	"context"
	"gorm.io/gorm"
)

type StudentRepository interface {
	Save(ctx context.Context, tx *gorm.DB, student domain.Student) domain.Student
	Update(ctx context.Context, tx *gorm.DB, student domain.Student) domain.Student
	Delete(ctx context.Context, tx *gorm.DB, student domain.Student)
	FindByNim(ctx context.Context, tx *gorm.DB, nim string) domain.Student
	FindAll(ctx context.Context, tx *gorm.DB) []domain.Student
}
