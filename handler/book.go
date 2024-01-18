package handler

import (
	"belajar-api/book"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *BookHandler {
	return &BookHandler{bookService}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *BookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": b,
	})
}

func (h *BookHandler) PostBookHandler(c *gin.Context) {
	//Mencoba menerima ada id dan title
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		//Cara menampilkan error
		var errorMessages []string

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) //Custom Error)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := h.bookService.CreateBook(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *BookHandler) UpdateBook(c *gin.Context){
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil{
		errorMessages := []string{}

		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field %s, condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := h.bookService.Update(id, bookRequest)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *BookHandler) DeleteBook(c *gin.Context){
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": b,
	})
}
