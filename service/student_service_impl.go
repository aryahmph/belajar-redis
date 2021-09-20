package service

import (
	"belajar-redis/helper"
	"belajar-redis/model/domain"
	"belajar-redis/model/web"
	"belajar-redis/repository"
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository, DB *gorm.DB, validate *validator.Validate) *StudentServiceImpl {
	return &StudentServiceImpl{StudentRepository: studentRepository, DB: DB, Validate: validate}
}

func (s *StudentServiceImpl) Create(ctx context.Context, request web.StudentCreateRequest) (web.StudentResponse, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return web.StudentResponse{}, err
	}

	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student := domain.Student{
		Name: request.Name,
	}
	student = s.StudentRepository.Save(ctx, tx, student)
	return helper.ToStudentResponse(student), nil
}

func (s *StudentServiceImpl) Update(ctx context.Context, request web.StudentUpdateRequest) (web.StudentResponse, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return web.StudentResponse{}, err
	}

	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	var studentResponse web.StudentResponse
	student, err := s.StudentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return studentResponse, err
	}
	return helper.ToStudentResponse(student), nil
}

func (s *StudentServiceImpl) Delete(ctx context.Context, studentId uint) error {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		return err
	}

	s.StudentRepository.Delete(ctx, tx, student)
	return nil
}

func (s *StudentServiceImpl) FindById(ctx context.Context, studentId uint) (web.StudentResponse, error) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		return web.StudentResponse{}, err
	}

	return helper.ToStudentResponse(student), nil
}

func (s *StudentServiceImpl) FindAll(ctx context.Context) []web.StudentResponse {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	students := s.StudentRepository.FindAll(ctx, tx)
	return helper.ToStudentResponses(students)
}
