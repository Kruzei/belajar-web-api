package handler

import (
	"belajar-api/app/borrowhistory/usecase"
	help "belajar-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BorrowHistoryHandler struct {
	borrowHistoryUsecase usecase.IBorrowHistoryUsecase
}

func NewBorrowHistoryHandler(borrowHistoryUsecase usecase.IBorrowHistoryUsecase) *BorrowHistoryHandler{
	return &BorrowHistoryHandler{borrowHistoryUsecase}
}

func (h *BorrowHistoryHandler) GetBorrowedBook(c *gin.Context){
	borrowedBook, errorObject := h.borrowHistoryUsecase.GetBorrowedBook()
	if errorObject != nil{
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, "failed to get borrowed book", errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "success to get all borrowed book", borrowedBook)
}

func (h *BorrowHistoryHandler) GetBorrowHistory(c *gin.Context){
	borrowHistories, errorObject := h.borrowHistoryUsecase.GetBorrowHistory()
	if errorObject != nil{
		errorObject := errorObject.(help.ErrorObject)
		help.FailedResponse(c, http.StatusBadRequest, errorObject.Message, errorObject.Err)
		return
	}

	help.SuccessResponse(c, http.StatusOK, "succes to get all book", borrowHistories)
}