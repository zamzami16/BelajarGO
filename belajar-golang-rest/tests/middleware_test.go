package tests

import (
	"belajar-go-rest/app"
	"belajar-go-rest/controller"
	"belajar-go-rest/exception"
	"belajar-go-rest/helper"
	"belajar-go-rest/middleware"
	"belajar-go-rest/model/domain"
	"belajar-go-rest/repository"
	"belajar-go-rest/service"
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func setupTestDb() *sql.DB {
	db, err := sql.Open("pgx", "user=axata password=axataposkenari host=localhost port=5433 dbname=belajar_go_rest_test sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func truncateCategories(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE category RESTART IDENTITY CASCADE")
	helper.PanicIfError(err)
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	router.PanicHandler = exception.ErrorHandler

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)
	router := setupRouter(db)

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", strings.NewReader(`{"name":"Category 1"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Category 1")
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)
	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", strings.NewReader(`{"name":""}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)
	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodPost, "/categories", strings.NewReader(`{"name":"Category 1"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func CreateCategory(db *sql.DB, ctx context.Context, name string) domain.Category {
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categoryRepository := repository.NewCategoryRepository()
	result := categoryRepository.Save(ctx, tx, domain.Category{
		Name: name,
	})

	return result
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	category := CreateCategory(db, context.Background(), "Category 1")

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), strings.NewReader(`{"name":"Category 1 Update"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	category := CreateCategory(db, context.Background(), "Category 1")

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), strings.NewReader(`{"name":""}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateCategoryNotFound(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:3000/api/categories/9999", strings.NewReader(`{"name":"Category 1 Update"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "RAHASIA")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	category := CreateCategory(db, context.Background(), "Category 1")

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteCategoryNotFound(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:3000/api/categories/9999", nil)
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestFindByIdSuccess(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	category := CreateCategory(db, context.Background(), "Category 1")

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), nil)
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestFindByIdNotFound(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories/9999", nil)
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestFindAllSuccess(t *testing.T) {
	db := setupTestDb()
	defer db.Close()
	truncateCategories(db)

	CreateCategory(db, context.Background(), "Category 1")
	CreateCategory(db, context.Background(), "Category 2")
	CreateCategory(db, context.Background(), "Category 3")
	router := setupRouter(db)
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	req.Header.Set("X-API-Key", "RAHASIA")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Category 1")
	assert.Contains(t, rec.Body.String(), "Category 2")
	assert.Contains(t, rec.Body.String(), "Category 3")
}
