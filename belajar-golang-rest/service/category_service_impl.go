package service

import (
	"belajar-go-rest/exception"
	"belajar-go-rest/helper"
	"belajar-go-rest/logging"
	"belajar-go-rest/middleware"
	"belajar-go-rest/model/domain"
	"belajar-go-rest/model/web"
	"belajar-go-rest/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	Logger             *logrus.Entry
}

func NewCategoryService(categoryRespository repository.CategoryRepository, db *sql.DB, validate *validator.Validate, loggerProvider *logging.LoggerProvider) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRespository,
		DB:                 db,
		Validate:           validate,
		Logger:             loggerProvider.GetLogger("CategoryService"),
	}
}

func (service *CategoryServiceImpl) getLoggerWithRequestID(ctx context.Context) *logrus.Entry {
	requestID := middleware.GetRequestID(ctx)
	return service.Logger.WithField("request_id", requestID)
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "create",
		"name":      request.Name,
	}).Info("Creating category")

	err := service.Validate.Struct(request)
	if err != nil {
		logger.WithField("error", err.Error()).Error("Validation failed")
		helper.PanicIfError(err)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	logger.WithFields(logrus.Fields{
		"id":   category.Id,
		"name": category.Name,
	}).Info("Category created successfully")

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "update",
		"id":        request.Id,
		"name":      request.Name,
	}).Info("Updating category")

	err := service.Validate.Struct(request)
	if err != nil {
		logger.WithField("error", err.Error()).Error("Validation failed")
		helper.PanicIfError(err)
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		// Use contextual error with request-aware logger
		panic(exception.NewContextualNotFoundError(err.Error(), logger))
	}

	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)
	logger.WithFields(logrus.Fields{
		"id":   category.Id,
		"name": category.Name,
	}).Info("Category updated successfully")

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "delete",
		"id":        categoryId,
	}).Info("Deleting category")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		// Use contextual error with request-aware logger
		panic(exception.NewContextualNotFoundError(err.Error(), logger))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
	logger.WithField("id", categoryId).Info("Category deleted successfully")
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "findById",
		"id":        categoryId,
	}).Info("Finding category by ID")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		// Use contextual error with request-aware logger
		panic(exception.NewContextualNotFoundError(err.Error(), logger))
	}

	logger.WithFields(logrus.Fields{
		"id":   category.Id,
		"name": category.Name,
	}).Info("Category found successfully")

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.Info("Finding all categories")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	logger.WithField("count", len(categories)).Info("Categories found successfully")

	return helper.ToCategoryResponses(categories)
}
