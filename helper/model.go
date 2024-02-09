package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
)

// konversi ke CategoryResponse
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// konversi ke CategoryResponses
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponse []web.CategoryResponse
	// Mengubah ke CategoryResponse lalu disimpan ke categoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}

	return categoryResponse

}
