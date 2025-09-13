package service

import (
	"belajar-go-rest/exception"
	"belajar-go-rest/helper"
	"belajar-go-rest/model/domain"
	"belajar-go-rest/model/web"
	"belajar-go-rest/repository"
	"context"
	"database/sql"
	"fmt"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryRespository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService  {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRespository,
		DB: db,
		Validate: validate,
	}
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	fmt.Println("masuk service")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if (err != nil){
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if (err != nil){
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if (err != nil){
		panic(exception.NewNotFoundError(err.Error()))
	}
	
	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}

