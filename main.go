package main

import (
	"belajar-api/app/book/handler"
	"belajar-api/app/book/repository"
	"belajar-api/app/book/usecase"
	"belajar-api/domain"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//Connect ke database
	dsn := "root:Himesama@tcp(127.0.0.1:3306)/belajar_web_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Connection Error")
	}

	db.AutoMigrate(domain.Books{})

	bookRepository := repository.NewRepository(db)
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
