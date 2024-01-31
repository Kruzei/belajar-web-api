package main

import (
	"belajar-api/app/book/handler"
	"belajar-api/app/book/repository"
	"belajar-api/app/book/usecase"
	borrowhistory_handler "belajar-api/app/borrowhistory/handler"
	borrowhistory_repository "belajar-api/app/borrowhistory/repository"
	borrowhistory_usecase "belajar-api/app/borrowhistory/usecase"
	user_handler "belajar-api/app/user/handler"
	user_repository "belajar-api/app/user/repository"
	user_usercase "belajar-api/app/user/usecase"
	"belajar-api/infrastructure"
	"belajar-api/infrastructure/database"
	"belajar-api/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connect ke database
	infrastructure.LoadEnv()

	database.ConnectDB()

	bookRepository := repository.NewRepository(database.DB)
	userRepository := user_repository.NewUserRepository(database.DB)
	borrowedHistoryRepository := borrowhistory_repository.NewBorrowHistoryRepository(database.DB)

	bookUsecase := usecase.NewBookUsecase(bookRepository)
	userUsecase := user_usercase.NewUserUsecase(userRepository)
	borrowedHistoryUsecase := borrowhistory_usecase.NewBorrowHistoryUsecase(borrowedHistoryRepository, bookRepository)

	bookHandler := handler.NewBookHandler(bookUsecase)
	userHandler := user_handler.NewUserHandler(userUsecase)
	borrowHistoryHandler := borrowhistory_handler.NewBorrowHistoryHandler(borrowedHistoryUsecase)

	router := rest.NewRest(gin.Default())

	router.RouteBooks(bookHandler)

	router.RouteUsers(userHandler)

	router.RouteBorrowHistories(borrowHistoryHandler)

	router.Run()
}
