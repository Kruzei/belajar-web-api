package handler

import (
	"belajar-api/helper"
	"belajar-api/app/book/usecase"
	"belajar-api/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase usecase.IBookUsecase
}

func NewBookHandler(bookUsecase usecase.IBookUsecase) *BookHandler {
	return &BookHandler{bookUsecase}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, errorObject := h.bookUsecase.FindAllBooks()
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "failed get all books", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes get all books", books)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, errorObject := h.bookUsecase.FindBookById(id)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "failed get book by id", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes get book by id", b)
}

func (h *BookHandler) PostBookHandler(c *gin.Context) {
	var bookRequest domain.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed bind book", err)
		return
	}

	book, errorObject := h.bookUsecase.CreateBook(bookRequest)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "failed to create", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes create book", book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var bookRequest domain.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "failed bind book", err)
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, errorObject := h.bookUsecase.Update(id, bookRequest)

	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "failed Update Book", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes update book", book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, errorObject := h.bookUsecase.Delete(id)

	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "failed delete book", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes delete book", b)
}
