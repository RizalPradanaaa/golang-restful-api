// go:build wireinject
//
//	+bulid wireinject
package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/repository"
	"golang-restful-api/service"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

func InitializedServer() *httprouter.Router {
	wire.Build(
		app.NewDb,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter,
	)
	return nil
}
