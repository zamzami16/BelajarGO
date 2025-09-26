package app

import (
	"belajar-go-rest/controller"
	"belajar-go-rest/exception"
	"belajar-go-rest/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler
	MapCategoryRoutes(router, categoryController)

	return router
}

func MapCategoryRoutes(router *httprouter.Router, categoryController controller.CategoryController) {
	router.GET("/api/categories", middleware.RequestIDMiddleware(categoryController.FindAll))
	router.GET("/api/categories/:categoryId", middleware.RequestIDMiddleware(categoryController.FindById))
	router.POST("/api/categories", middleware.RequestIDMiddleware(categoryController.Create))
	router.PUT("/api/categories/:categoryId", middleware.RequestIDMiddleware(categoryController.Update))
	router.DELETE("/api/categories/:categoryId", middleware.RequestIDMiddleware(categoryController.Delete))
}
