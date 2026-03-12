//go:build wireinject
// +build wireinject

package main

import (
	"belajar-go-rest/app"
	"belajar-go-rest/controller"
	"belajar-go-rest/logging"
	"belajar-go-rest/middleware"
	"belajar-go-rest/repository"
	"belajar-go-rest/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)

// Providers for pgxpool-based implementation
var categoryPgxSet = wire.NewSet(
	repository.NewCategoryRepository, // Use existing repository
	service.NewCategoryServicePgx,    // Use pgxpool service with existing repository
	controller.NewCategoryController,
)

// Original SQL-based providers (kept for compatibility)
var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewPgxPool, // Use pgxpool instead of NewDB
		validator.New,
		logging.NewLoggerProvider,
		categoryPgxSet, // Use pgxpool-based implementation
		app.NewRouter,
		middleware.NewAuthMiddleware,
		newServer,
	)
	return nil
}

// Alternative initializer with SQL-based implementation
func InitializeServerSQL() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		logging.NewLoggerProvider,
		categorySet,
		app.NewRouter,
		middleware.NewAuthMiddleware,
		newServer,
	)
	return nil
}

// Helper to create http.Server
func newServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
