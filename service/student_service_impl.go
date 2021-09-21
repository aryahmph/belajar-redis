package service

import (
	"belajar-redis/helper"
	"belajar-redis/model/domain"
	"belajar-redis/model/web"
	"belajar-redis/repository"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type StudentServiceImpl struct {
	StudentRepository repository.StudentRepository
	StudentCache      repository.StudentCache
	DB                *gorm.DB
	Redis             *redis.Client
	Validate          *validator.Validate
}

func NewStudentServiceImpl(studentRepository repository.StudentRepository, studentCache repository.StudentCache, DB *gorm.DB, redis *redis.Client, validate *validator.Validate) *StudentServiceImpl {
	return &StudentServiceImpl{StudentRepository: studentRepository, StudentCache: studentCache, DB: DB, Redis: redis, Validate: validate}
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
	// Db
	student = s.StudentRepository.Save(ctx, tx, student)
	studentResponse := helper.ToStudentResponse(student)

	// Cache
	s.StudentCache.Delete(ctx, s.Redis, studentResponse.Id)
	s.StudentCache.SetById(ctx, s.Redis, studentResponse)

	return studentResponse, nil
}

func (s *StudentServiceImpl) Update(ctx context.Context, request web.StudentUpdateRequest) (web.StudentResponse, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		return web.StudentResponse{}, err
	}

	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.StudentResponse{}, err
	}
	student.Name = request.Name

	student = s.StudentRepository.Update(ctx, tx, student)
	studentResponse := helper.ToStudentResponse(student)

	s.StudentCache.Delete(ctx, s.Redis, student.ID)
	s.StudentCache.SetById(ctx, s.Redis, studentResponse)
	return studentResponse, nil
}

func (s *StudentServiceImpl) Delete(ctx context.Context, studentId uint) error {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	student, err := s.StudentRepository.FindById(ctx, tx, studentId)
	if err != nil {
		return err
	}

	s.StudentRepository.Delete(ctx, tx, student)
	//s.StudentCache.Delete(ctx, s.Redis, studentId)
	return nil
}

func (s *StudentServiceImpl) FindById(ctx context.Context, studentId uint) (web.StudentResponse, error) {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	studentResponse, err := s.StudentCache.FindById(ctx, s.Redis, studentId)
	if err == redis.Nil {
		student, err := s.StudentRepository.FindById(ctx, tx, studentId)
		if err != nil {
			return web.StudentResponse{}, err
		}
		return helper.ToStudentResponse(student), nil
	} else if err != nil {
		panic(err)
	} else {
		return studentResponse, nil
	}
}

func (s *StudentServiceImpl) FindAll(ctx context.Context) []web.StudentResponse {
	tx := s.DB.Begin()
	defer helper.CommitOrRollback(tx)

	studentResponses, err := s.StudentCache.FindAll(ctx, s.Redis)
	if err == redis.Nil {
		students := s.StudentRepository.FindAll(ctx, tx)
		studentResponses = helper.ToStudentResponses(students)
		s.StudentCache.SetAll(ctx, s.Redis, studentResponses)
		return studentResponses
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(studentResponses)
		return studentResponses
	}
}
