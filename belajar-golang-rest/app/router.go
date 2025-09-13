package app

import (
	"belajar-go-rest/controller"

	"github.com/julienschmidt/httprouter"
)

func AddCategoryRoutes(router *httprouter.Router, categoryController controller.CategoryController) {
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}
