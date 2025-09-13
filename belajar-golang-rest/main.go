package main

import (
	"belajar-go-rest/app"
	"belajar-go-rest/controller"
	"belajar-go-rest/exception"
	"belajar-go-rest/helper"
	"belajar-go-rest/middleware"
	"belajar-go-rest/repository"
	"belajar-go-rest/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRespository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRespository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	app.AddCategoryRoutes(router, categoryController)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	fmt.Println("Berjalan di local host:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
