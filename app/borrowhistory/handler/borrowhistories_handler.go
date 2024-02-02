package handler

import (
	"belajar-api/app/borrowhistory/usecase"
	help "belajar-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BorrowHistoryHandler struct {
	borrowHistoryUsecase usecase.IBorrowHistoryUsecase
}

func NewBorrowHistoryHandler(borrowHistoryUsecase usecase.IBorrowHistoryUsecase) *BorrowHistoryHandler {
	return &BorrowHistoryHandler{borrowHistoryUsecase}
}

func (h *BorrowHistoryHandler) GetBorrowedBook(c *gin.Context) {
	borrowedBook, errorObject := h.borrowHistoryUsecase.GetBorrowedBook()
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "success to get all borrowed book", borrowedBook)
}

func (h *BorrowHistoryHandler) GetBorrowHistory(c *gin.Context) {
	borrowHistories, errorObject := h.borrowHistoryUsecase.GetBorrowHistory()
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes to get all book", borrowHistories)
}

func (h *BorrowHistoryHandler) BorrowBook(c *gin.Context) {
	bookIdString := c.Param("bookid")
	bookId, _ := strconv.Atoi(bookIdString)

	borrowedBook, errorObject := h.borrowHistoryUsecase.BorrowBook(bookId, c)
	if errorObject != nil {
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "success borrow book", borrowedBook)
}

func (h *BorrowHistoryHandler) ReturnBook(c *gin.Context){
	bookIdString := c.Param("bookid")
	bookId, _ := strconv.Atoi(bookIdString)

	returnBook, errorObject := h.borrowHistoryUsecase.ReturnBook(bookId, c)
	if errorObject != nil{
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "success return book", returnBook)
}
