package controller

import (
	"belajar-redis/exception"
	"belajar-redis/helper"
	"belajar-redis/model/web"
	"belajar-redis/service"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type StudentControllerImpl struct {
	StudentService service.StudentService
}

func NewStudentControllerImpl(studentService service.StudentService) *StudentControllerImpl {
	return &StudentControllerImpl{StudentService: studentService}
}

func (s *StudentControllerImpl) Create(writer http.ResponseWriter, request *http.Request) {
	studentCreateRequest := web.StudentCreateRequest{}
	err := helper.ReadFromRequestBody(request, &studentCreateRequest)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	studentResponse, err := s.StudentService.Create(request.Context(), studentCreateRequest)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s *StudentControllerImpl) Update(writer http.ResponseWriter, request *http.Request) {
	studentUpdateRequest := web.StudentUpdateRequest{}
	err := helper.ReadFromRequestBody(request, &studentUpdateRequest)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	studentId := chi.URLParam(request, "studentId")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	studentUpdateRequest.Id = uint(id)
	studentResponse, err := s.StudentService.Update(request.Context(), studentUpdateRequest)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s *StudentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request) {
	studentId := chi.URLParam(request, "studentId")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	err = s.StudentService.Delete(request.Context(), uint(id))
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s *StudentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request) {
	studentId := chi.URLParam(request, "studentId")
	id, err := strconv.Atoi(studentId)
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}

	studentResponse, err := s.StudentService.FindById(request.Context(), uint(id))
	if err != nil {
		exception.ErrorHandler(writer, request, err)
		return
	}
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s *StudentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request) {
	studentResponses := s.StudentService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   studentResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
