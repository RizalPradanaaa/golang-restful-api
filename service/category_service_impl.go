package service

import (
	"context"
	"database/sql"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// Database Transaksional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Ambil data dari request
	category := domain.Category{
		Name: request.Name,
	}

	service.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// Database Transaksional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Cek data id dahulu
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// set data name
	category.Name = request.Name

	service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// Database Transaksional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Cek data id dahulu
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categotyId int) web.CategoryResponse {
	// Database Transaksional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categotyId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// Database Transaksional
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
