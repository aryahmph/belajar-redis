package repository

import (
	"belajar-redis/helper"
	"belajar-redis/model/domain"
	"context"
	"gorm.io/gorm"
)

type StudentRepositoryImpl struct {
}

func NewStudentRepositoryImpl() *StudentRepositoryImpl {
	return &StudentRepositoryImpl{}
}

func (s *StudentRepositoryImpl) Save(ctx context.Context, tx *gorm.DB, student domain.Student) domain.Student {
	err := tx.WithContext(ctx).Create(&student).Error
	helper.PanicIfError(err)
	return student
}

func (s *StudentRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, student domain.Student) domain.Student {
	err := tx.WithContext(ctx).Save(&student).Error
	helper.PanicIfError(err)
	return student
}

func (s *StudentRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, student domain.Student) {
	err := tx.WithContext(ctx).Delete(&student).Error
	helper.PanicIfError(err)
}

func (s *StudentRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, studentId uint) (domain.Student, error) {
	var student domain.Student
	err := tx.WithContext(ctx).Where("id = ?", studentId).First(&student).Error
	if err != nil {
		return student, err
	} else {
		return student, nil
	}
}

func (s *StudentRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) []domain.Student {
	var students []domain.Student
	err := tx.WithContext(ctx).Find(&students).Error
	helper.PanicIfError(err)
	return students
}
