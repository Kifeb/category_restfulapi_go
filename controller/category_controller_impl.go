package controller

import (
	"belajar-restful-api/helper"
	"belajar-restful-api/model/web"
	"belajar-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryCreateReq := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateReq)

	categoryRes := controller.CategoryService.Create(r.Context(), categoryCreateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRes,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryUpdateReq := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateReq)

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErr(err)

	categoryUpdateReq.Id = id

	categoryRes := controller.CategoryService.Update(r.Context(), categoryUpdateReq)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRes,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErr(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfErr(err)

	categoryRes := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryRes,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
