package main

import (
	"belajar-api/app/book/handler"
	"belajar-api/app/book/repository"
	"belajar-api/app/book/usecase"
	user_handler "belajar-api/app/user/handler"
	user_repository "belajar-api/app/user/repository"
	user_usercase "belajar-api/app/user/usecase"
	"belajar-api/infrastructure"
	"belajar-api/infrastructure/database"
	"belajar-api/middleware"

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

	validate := middleware.RequireAuth

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books/:id", validate, bookHandler.GetBook)
	v1.GET("/books", validate, bookHandler.GetBooks)
	v1.POST("/books", validate, bookHandler.PostBookHandler)
	v1.PUT("/books/:id", validate, bookHandler.UpdateBook)
	v1.DELETE("books/:id", validate, bookHandler.DeleteBook)

	v1.GET("/users", validate ,userHandler.FindAllUsers)
	v1.GET("/users/:id", validate, userHandler.FindUser)
	v1.POST("/signup", userHandler.SignUp)
	v1.POST("/login", userHandler.Login)
	v1.PUT("/users/:id", validate, userHandler.UpdateUser)
	v1.DELETE("/users/:id", validate, userHandler.DeleteUser)

	router.Run() //Default portnya localhost:8080
}
