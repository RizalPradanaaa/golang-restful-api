package main

import (
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// db := app.NewDb()
	// validate := validator.New()

	// categoryRepository := repository.NewCategoryRepository()
	// categoryService := service.NewCategoryService(categoryRepository, db, validate)
	// categoryController := controller.NewCategoryController(categoryService)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(InitializedServer()),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
