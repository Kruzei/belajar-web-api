package usecase

import (
	"belajar-api/app/book/repository"
	// borrowRepository "belajar-api/app/borrowhistories/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type IBookUsecase interface {
	GetAllBooks() ([]domain.Books, any)
	GetBookById(id int) (domain.Books, any)
	GetAvailableBook() ([]domain.Books, any)
	CreateBook(bookRequest domain.BookRequest) (domain.Books, any)
	Update(id int, bookRequest domain.BookRequest) (domain.Books, any)
	Delete(id int) (domain.Books, any)
	BorrowBook(bookId int, c *gin.Context) (domain.Books, any)
	ReturnBook(bookId int, c *gin.Context) (domain.Books, any)
}

type BookUsecase struct {
	bookRepository repository.IBookRepository
}

func NewBookUsecase(repository repository.IBookRepository) *BookUsecase {
	return &BookUsecase{repository}
}

func (s *BookUsecase) GetAllBooks() ([]domain.Books, any) {
	var books []domain.Books
	err := s.bookRepository.GetAllBooks(&books)
	if err != nil {
		return books, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "error occured when find all book",
			Err:     err,
		}
	}

	return books, err
}

func (s *BookUsecase) GetBookById(id int) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.GetBookById(&book, id)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "id is not exist",
			Err:     errors.New("id not found"),
		}
	}

	return book, err
}

func (s *BookUsecase) GetAvailableBook() ([]domain.Books, any) {
	var books []domain.Books
	err := s.bookRepository.GetAvailableBook(&books)
	if err != nil {
		return books, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "error occured when find available book",
			Err:     err,
		}
	}

	return books, err
}

func (s *BookUsecase) CreateBook(bookRequest domain.BookRequest) (domain.Books, any) {
	book := domain.Books{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Status:      bookRequest.Status,
	}

	err := s.bookRepository.CreateBook(&book)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to make book data",
			Err:     err,
		}
	}

	return book, nil
}

func (s *BookUsecase) Update(id int, bookRequest domain.BookRequest) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.GetBookById(&book, id)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "id is not exist",
			Err:     errors.New("id not found"),
		}
	}

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description

	err = s.bookRepository.Update(&book)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to update book data",
			Err:     err,
		}
	}
	return book, err
}

func (s *BookUsecase) Delete(id int) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.GetBookById(&book, id)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "id is not exist",
			Err:     errors.New("id not found"),
		}
	}

	err = s.bookRepository.Delete(&book)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to delete book",
			Err:     err,
		}
	}

	return book, err
}

func (s *BookUsecase) BorrowBook(bookId int, c *gin.Context) (domain.Books, any) {
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
	err := s.bookRepository.GetBookById(&book, bookId)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "book is not exist",
			Err:     err,
		}
	}

	if book.Status != "available" {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusBadRequest,
			Message: "book not available",
			Err:     errors.New(""),
		}
	}

	book.Status = "NOT AVAILABLE"
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

	err = s.bookRepository.Borrow(&borrowBook)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to borrow book",
			Err:     err,
		}
	}

	return book, err
}

func (s *BookUsecase) ReturnBook(bookId int, c *gin.Context) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.GetBookById(&book, bookId)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "book is not exist",
			Err:     err,
		}
	}

	var borrowedBook domain.BorrowHistories
	err = s.bookRepository.GetBorrowedBook(&borrowedBook, bookId)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusNotFound,
			Message: "borrowed book not found",
			Err:     err,
		}
	}

	book.Status = "AVAILABLE"
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

	err = s.bookRepository.ReturnBook(&borrowedBook)
	if err != nil {
		return domain.Books{}, help.ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: "failed to return book",
			Err:     err,
		}
	}

	return book, err
}
