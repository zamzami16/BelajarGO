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

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

type CategoryServicePgxImpl struct {
	CategoryRepository repository.CategoryRepository // Use existing interface
	PgxPool            *pgxpool.Pool
	Validate           *validator.Validate
	Logger             *logrus.Entry
}

func NewCategoryServicePgx(categoryRepository repository.CategoryRepository, pgxPool *pgxpool.Pool, validate *validator.Validate, loggerProvider *logging.LoggerProvider) CategoryService {
	return &CategoryServicePgxImpl{
		CategoryRepository: categoryRepository, // Use existing repository
		PgxPool:            pgxPool,
		Validate:           validate,
		Logger:             loggerProvider.GetLogger("CategoryServicePgx"),
	}
}

// Helper function to get logger with request ID
func (service *CategoryServicePgxImpl) getLoggerWithRequestID(ctx context.Context) *logrus.Entry {
	requestID := middleware.GetRequestID(ctx)
	return service.Logger.WithField("request_id", requestID)
}

func (service *CategoryServicePgxImpl) Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "create",
		"name":      request.Name,
	}).Info("Creating category with pgxpool transaction")

	err := service.Validate.Struct(request)
	if err != nil {
		logger.WithField("error", err.Error()).Error("Validation failed")
		helper.PanicIfError(err)
	}

	// Get sql.DB from pgxpool using stdlib connector
	sqlDB := stdlib.OpenDBFromPool(service.PgxPool)

	// Use existing transaction pattern with sql.Tx
	tx, err := sqlDB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	// Use existing repository interface with sql.Tx
	category = service.CategoryRepository.Save(ctx, tx, category)

	logger.WithFields(logrus.Fields{
		"id":   category.Id,
		"name": category.Name,
	}).Info("Category created successfully with pgxpool transaction")

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServicePgxImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "update",
		"id":        request.Id,
		"name":      request.Name,
	}).Info("Updating category with pgxpool transaction")

	err := service.Validate.Struct(request)
	if err != nil {
		logger.WithField("error", err.Error()).Error("Validation failed")
		helper.PanicIfError(err)
	}

	// Get sql.DB from pgxpool using stdlib connector
	sqlDB := stdlib.OpenDBFromPool(service.PgxPool)

	// Use existing transaction pattern with sql.Tx
	tx, err := sqlDB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Check if category exists first
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewContextualNotFoundError(err.Error(), logger))
	}

	// Update category
	category.Name = request.Name
	category = service.CategoryRepository.Update(ctx, tx, category)

	logger.WithFields(logrus.Fields{
		"id":   category.Id,
		"name": category.Name,
	}).Info("Category updated successfully with pgxpool transaction")

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServicePgxImpl) Delete(ctx context.Context, categoryId int) {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "delete",
		"id":        categoryId,
	}).Info("Deleting category with pgxpool transaction")

	// Get sql.DB from pgxpool using stdlib connector
	sqlDB := stdlib.OpenDBFromPool(service.PgxPool)

	// Use existing transaction pattern with sql.Tx
	tx, err := sqlDB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Check if category exists first
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewContextualNotFoundError(err.Error(), logger))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
	logger.WithField("id", categoryId).Info("Category deleted successfully with pgxpool transaction")
}

func (service *CategoryServicePgxImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.WithFields(logrus.Fields{
		"operation": "findById",
		"id":        categoryId,
	}).Info("Finding category by ID with pgxpool")

	// Get sql.DB from pgxpool using stdlib connector
	sqlDB := stdlib.OpenDBFromPool(service.PgxPool)

	// Read operations can use transaction for consistency but it's optional
	tx, err := sqlDB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewContextualNotFoundError(err.Error(), logger))
	}

	logger.WithFields(logrus.Fields{
		"id":   category.Id,
		"name": category.Name,
	}).Info("Category found successfully with pgxpool")

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServicePgxImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	logger := service.getLoggerWithRequestID(ctx)

	logger.Info("Finding all categories with pgxpool")

	// Get sql.DB from pgxpool using stdlib connector
	sqlDB := stdlib.OpenDBFromPool(service.PgxPool)

	// Read operations can use transaction for consistency
	tx, err := sqlDB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)
	logger.WithField("count", len(categories)).Info("Categories found successfully with pgxpool")

	return helper.ToCategoryResponses(categories)
}
