package usecase

import (
	bookRepository "belajar-api/app/book/repository"
	"belajar-api/app/borrowhistory/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
)

type IBorrowHistoryUsecase interface {
	GetBorrowedBook() ([]domain.BorrowedBookResponse, any)
	GetBorrowHistory() ([]domain.BorrowHistoryResponse, any)
}

type BorrowHistoryUsecase struct {
	borrowHistoryRepository repository.IBorrowHistory
	bookRepository          bookRepository.IBookRepository
}

func NewBorrowHistoryUsecase(repository repository.IBorrowHistory, bookRepository bookRepository.IBookRepository) *BorrowHistoryUsecase {
	return &BorrowHistoryUsecase{repository, bookRepository}
}

func (s *BorrowHistoryUsecase) GetBorrowedBook() ([]domain.BorrowedBookResponse, any) {
	var books []domain.Books
	err := s.bookRepository.GetBookByCondition(&books, "status = ?", "NOT AVAIBLE")
	if err != nil {
		return []domain.BorrowedBookResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to get all borrowed book",
			Err:     err,
		}
	}

	var user []domain.Users
	err2 := s.borrowHistoryRepository.GetBorrowedBook(&books, &user)
	if err2 != nil {
		return []domain.BorrowedBookResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "book not found",
			Err:     errors.New(""),
		}
	}

	var booksResponse []domain.BorrowedBookResponse
	for i, b := range books {
		bookResponse := domain.BorrowedBookResponse{
			Name:        user[i].Name,
			Title:       b.Title,
			Description: b.Description,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	return booksResponse, nil
}

func (s *BorrowHistoryUsecase) GetBorrowHistory() ([]domain.BorrowHistoryResponse, any) {
	var borrowHistories []domain.BorrowHistories
	err := s.borrowHistoryRepository.GetBorrowHistory(&borrowHistories)
	if err != nil {
		return []domain.BorrowHistoryResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "error occured when find all book histories",
			Err:     err,
		}
	}

	var books []domain.Books
	err = s.borrowHistoryRepository.GetAllBookById(&books)
	if err != nil {
		return []domain.BorrowHistoryResponse{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "DISINI",
			Err:     err,
		}
	}
	var borrowHistoriesResponse []domain.BorrowHistoryResponse

	for i, b := range borrowHistories {
		borrowHistoryResponse := domain.BorrowHistoryResponse{
			Title:      books[i].Title,
			BorrowTime: b.BorrowTime,
			ReturnTime: b.ReturnTime,
		}
		borrowHistoriesResponse = append(borrowHistoriesResponse, borrowHistoryResponse)
	}

	return borrowHistoriesResponse, nil
}
