//go:build wireinject
// +build wireinject

package main

import (
	"belajar-go-rest/app"
	"belajar-go-rest/controller"
	"belajar-go-rest/middleware"
	"belajar-go-rest/repository"
	"belajar-go-rest/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
)

// Providers
var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
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
