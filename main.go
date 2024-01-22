package main

import (
	"belajar-api/app/book/handler"
	"belajar-api/app/book/repository"
	"belajar-api/app/book/usecase"
	"belajar-api/infrastructure"
	"belajar-api/infrastructure/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connect ke database
	infrastructure.LoadEnv()

	database.ConnectDB()

	bookRepository := repository.NewRepository(database.DB)
	bookUsecase := usecase.NewBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookUsecase)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", bookHandler.GetBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("books/:id", bookHandler.DeleteBook)

	router.Run() //Default portnya localhost:8080, kalau mau di custom bisa kayak gini router.Run(":8888")
}
