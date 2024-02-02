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
	v1.POST("/users/books/:bookid/borrows", validate, bookHandler.BorrowBook)
	v1.POST("/users/books/:bookid/returns", validate, bookHandler.ReturnBook)
}

func (rest *Rest) RouteUsers(userHandler *user_handler.UserHandler) {
	validate := middleware.RequireAuth
	authorization := middleware.OnlyAdmin
	v1 := rest.gin.Group("/v1")

	v1.GET("/users", validate, authorization, userHandler.FindAllUsers)
	v1.GET("/users/:id", validate, authorization, userHandler.FindUser)
	v1.POST("/signup", userHandler.SignUp)
	v1.POST("/login", userHandler.Login)
	v1.PUT("/users/:id", validate, userHandler.UpdateUser)
	v1.DELETE("/users/:id", validate, userHandler.DeleteUser)
}

func (rest *Rest) RouteBorrowHistories(borrowHistoryHandler *borrowhistory_handler.BorrowHistoryHandler) {
	validate := middleware.RequireAuth
	authorization := middleware.OnlyAdmin
	v1 := rest.gin.Group("/v1")

	v1.GET("/books/borrowed", validate, authorization, borrowHistoryHandler.GetBorrowedBook)
	v1.GET("/books/books-history", validate, authorization, borrowHistoryHandler.GetBorrowHistory)
}

func (rest *Rest) Run() {
	rest.gin.Run()
}
