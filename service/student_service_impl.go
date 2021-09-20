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
		Nim:  request.Nim,
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

	var studentResponses web.StudentResponse
	student, err := s.StudentRepository.FindByNim(ctx, tx, request.Nim)
	if err != nil {
		return studentResponses, err
	}
	return helper.ToStudentResponse(student), nil
}

func (s *StudentServiceImpl) Delete(ctx context.Context, studentNim string) error {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindByNim(ctx, tx, studentNim)
	if err != nil {
		return err
	}

	s.StudentRepository.Delete(ctx, tx, student)
	return nil
}

func (s *StudentServiceImpl) FindByNim(ctx context.Context, studentNim string) (web.StudentResponse, error) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindByNim(ctx, tx, studentNim)
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
