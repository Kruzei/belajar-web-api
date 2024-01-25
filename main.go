package main

import (
	"belajar-api/app/book/handler"
	"belajar-api/app/book/repository"
	user_repository"belajar-api/app/user/repository"
	user_usercase"belajar-api/app/user/usecase"
	user_handler"belajar-api/app/user/handler"
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
	userRepository := user_repository.NewUserRepository(database.DB)

	bookUsecase := usecase.NewBookUsecase(bookRepository)
	userUsecase := user_usercase.NewUserUsecase(userRepository)

	bookHandler := handler.NewBookHandler(bookUsecase)
	userHandler := user_handler.NewUserHandler(userUsecase)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", bookHandler.GetBook)
	v1.GET("/books", bookHandler.GetBooks)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("books/:id", bookHandler.DeleteBook)

	v1.GET("/users", userHandler.FindAllUsers)
	v1.GET("/users/:id", userHandler.FindUser)
	v1.POST("/users", userHandler.SignUp)
	v1.POST("/login", userHandler.Login)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.DELETE("/users/:id", userHandler.DeleteUser)

	router.Run() //Default portnya localhost:8080
}
