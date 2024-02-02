package usecase

import (
	bookRepository "belajar-api/app/book/repository"
	"belajar-api/app/borrowhistory/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
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
	var borrowHistories []domain.BorrowHistories
	err := s.borrowHistoryRepository.GetBorrowedBooks(&borrowHistories)
	if err != nil {
		return []domain.BorrowedBookResponse{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "book not found",
			Err:     err,
		}
	}

	var booksResponse []domain.BorrowedBookResponse
	for _, b := range borrowHistories {
		bookResponse := domain.BorrowedBookResponse{
			Name:        b.User.Name,
			Title:       b.Book.Title,
			Description: b.Book.Title,
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

	var borrowHistoriesResponse []domain.BorrowHistoryResponse

	for _, b := range borrowHistories {
		borrowHistoryResponse := domain.BorrowHistoryResponse{
			Title:      b.Book.Title,
			BorrowTime: b.BorrowTime,
			ReturnTime: b.ReturnTime,
		}
		borrowHistoriesResponse = append(borrowHistoriesResponse, borrowHistoryResponse)
	}

	return borrowHistoriesResponse, nil
}
