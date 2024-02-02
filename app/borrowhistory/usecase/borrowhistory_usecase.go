package usecase

import (
	bookRepository "belajar-api/app/book/repository"
	"belajar-api/app/borrowhistory/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IBorrowHistoryUsecase interface {
	GetBorrowedBook() ([]domain.BorrowedBookResponse, any)
	GetBorrowHistory() ([]domain.BorrowHistoryResponse, any)
	BorrowBook(bookId int, c *gin.Context) (domain.Books, any)
	ReturnBook(bookId int, c *gin.Context) (domain.Books, any)
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

func (s *BorrowHistoryUsecase) BorrowBook(bookId int, c *gin.Context) (domain.Books, any) {
	user, exists := c.Get("user")
	if !exists {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "user not found",
			Err:     errors.New(""),
		}
	}

	userId := user.(domain.Users).ID
	var book domain.Books
	err := s.bookRepository.FindById(&book, bookId)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "book is not exist",
			Err:     err,
		}
	}

	if book.Status != "AVAIBLE" {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "book not avaible",
			Err:     errors.New(""),
		}
	}

	book.Status = "NOT AVAIBLE"
	err = s.bookRepository.Update(&book)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to borrow book",
			Err:     err,
		}
	}

	date := time.Now()

	var borrowBook = domain.BorrowHistories{
		UserId:     userId,
		BookId:     book.ID,
		BorrowTime: date,
	}

	err = s.borrowHistoryRepository.Borrow(&borrowBook)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to borrow book",
			Err:     err,
		}
	}

	return book, err
}

func (s *BorrowHistoryUsecase) ReturnBook(bookId int, c *gin.Context) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.FindById(&book, bookId)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "book is not exist",
			Err:     err,
		}
	}

	var borrowedBook domain.BorrowHistories
	err = s.bookRepository.FindBorrowedBook(&borrowedBook, bookId)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code: http.StatusNotFound,
			Message: "borrowed book not found",
			Err: err,
		}
	}

	book.Status = "AVAIBLE"
	err = s.bookRepository.Update(&book)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to return book",
			Err:     err,
		}
	}

	date := time.Now()
	borrowedBook.ReturnTime = date

	err = s.borrowHistoryRepository.ReturnBook(&borrowedBook)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to return book",
			Err:     err,
		}
	}

	return book, err
}