package helper

import (
	"belajar-redis/model/domain"
	"belajar-redis/model/web"
)

func ToStudentResponse(student domain.Student) web.StudentResponse {
	return web.StudentResponse{
		Id:   student.ID,
		Name: student.Name,
		Nim:  student.Nim,
	}
}

func ToStudentResponses(students []domain.Student) []web.StudentResponse {
	var studentResponses []web.StudentResponse
	for _, student := range students {
		studentResponses = append(studentResponses, ToStudentResponse(student))
	}
	return studentResponses
}
