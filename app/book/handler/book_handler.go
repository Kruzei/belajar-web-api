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
	books, err := h.bookUsecase.FindAllBooks()
	if err != nil {
		err := err.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed get all books", err.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Succes get all books", books)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookUsecase.FindBookById(id)
	if err != nil {
		err := err.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed get book by id", err.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Succes get book by id", b)
}

func (h *BookHandler) PostBookHandler(c *gin.Context) {
	//Mencoba menerima ada id dan title
	var bookRequest domain.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		//Cara menampilkan error
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind book", err)
		return
	}

	book, err2 := h.bookUsecase.CreateBook(bookRequest)
	if err2 != nil {
		err2 := err2.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed to create", err2.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Succes create book", book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var bookRequest domain.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		help.FailedResponse(c, http.StatusBadRequest, "Failed bind book", err)
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err2 := h.bookUsecase.Update(id, bookRequest)

	if err != nil {
		err2 := err2.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed Update Book", err2.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Succes update book", book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookUsecase.Delete(id)

	if err != nil {
		err := err.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "Failed delete book", err.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "Succes delete book", b)
}
