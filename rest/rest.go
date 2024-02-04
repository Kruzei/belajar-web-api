package rest

import (
	book_handler "belajar-api/app/book/handler"
	borrowhistory_handler "belajar-api/app/borrowhistory/handler"
	user_handler "belajar-api/app/user/handler"
	"belajar-api/middleware"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	gin *gin.Engine
}

func NewRest(gin *gin.Engine) Rest {
	return Rest{
		gin: gin,
	}
}

func (rest *Rest) RouteBooks(bookHandler *book_handler.BookHandler) {
	validate := middleware.RequireAuth
	authorization := middleware.OnlyAdmin
	v1 := rest.gin.Group("/v1")

	v1.GET("/books/:id", validate, bookHandler.GetBook)
	v1.GET("/books", validate, bookHandler.GetBooks)
	v1.POST("/books", validate, authorization, bookHandler.CreateBook)
	v1.PUT("/books/:id", validate, authorization, bookHandler.UpdateBook)
	v1.DELETE("books/:id", validate, authorization, bookHandler.DeleteBook)

	v1.GET("users/books/available", validate, bookHandler.GetAvailableBook)
}

func (rest *Rest) RouteUsers(userHandler *user_handler.UserHandler) {
	validate := middleware.RequireAuth
	authorization := middleware.OnlyAdmin
	v1 := rest.gin.Group("/v1")

	v1.GET("/users", validate, authorization, userHandler.GetAllUsers)
	v1.GET("/users/:id", validate, authorization, userHandler.GetUser)
	v1.POST("/signup", userHandler.SignUp)
	v1.POST("/login", userHandler.Login)
	v1.PUT("/users/:id", validate, userHandler.UpdateUser)
	v1.DELETE("/users/:id", validate, userHandler.DeleteUser)
	// v1.GET("/user/:user_id", userHandler.GetUser)
}

func (rest *Rest) RouteBorrowHistories(borrowHistoryHandler *borrowhistory_handler.BorrowHistoryHandler) {
	validate := middleware.RequireAuth
	authorization := middleware.OnlyAdmin
	v1 := rest.gin.Group("/v1")

	v1.GET("/books/borrowed", validate, authorization, borrowHistoryHandler.GetBorrowedBook)
	v1.GET("/books/books-history", validate, authorization, borrowHistoryHandler.GetBorrowHistory)
	v1.GET("/users/books/borrow-histories", validate, borrowHistoryHandler.GetUserBorrowedBook)
	v1.POST("users/books/:book_id/borrow-histories", validate, borrowHistoryHandler.BorrowBook)
	v1.PUT("users/books/borrow-histories/:borrow-histories-id", validate, borrowHistoryHandler.ReturnBook)
}

func (rest *Rest) Run() {
	rest.gin.Run()
}
