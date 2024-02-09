package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// membaca request body
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Mengubah ke response json
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// membaca request body
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	// Ambil Id dari paramater
	categoryId := params.ByName("categoryId")
	// Konversi ke int
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Mengubah response ke json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil id dari param
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	// Mengubah response ke json
	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Ambil id dari param
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	// Mengubah response ke json
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	// Mengubah response ke json
	helper.WriteToResponseBody(writer, webResponse)
}
