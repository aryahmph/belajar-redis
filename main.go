package main

import (
	"belajar-redis/app"
	"belajar-redis/controller"
	"belajar-redis/helper"
	"belajar-redis/model/domain"
	"belajar-redis/repository"
	"belajar-redis/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db := app.NewDB()
	redis := app.NewRedis()
	validate := validator.New()
	err := db.AutoMigrate(&domain.Student{})
	helper.PanicIfError(err)

	studentRepository := repository.NewStudentRepositoryImpl()
	studentCache := repository.NewStudentCacheImpl()
	studentService := service.NewStudentServiceImpl(studentRepository, studentCache, db, redis, validate)
	studentController := controller.NewStudentControllerImpl(studentService)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api/v1/students", func(router chi.Router) {
		router.Post("/", studentController.Create)
		router.Get("/", studentController.FindAll)
		router.Get("/{studentId}", studentController.FindById)
		router.Put("/{studentId}", studentController.Update)
		router.Delete("/{studentId}", studentController.Delete)
	})

	server := http.Server{
		Handler: router,
		Addr:    ":3000",
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
