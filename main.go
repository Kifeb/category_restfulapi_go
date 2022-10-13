package main

import (
	"belajar-restful-api/app"
	"belajar-restful-api/controller"
	"belajar-restful-api/exception"
	"belajar-restful-api/helper"
	"belajar-restful-api/middlewere"
	"belajar-restful-api/repository"
	"belajar-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: middlewere.NewAuthMiddlewere(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
