package service

import (
	"context"
	"golang-restful-api/model/web"
)

type CategoryService interface {
	Save(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(cxt context.Context, catgoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
