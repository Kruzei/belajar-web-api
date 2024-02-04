package usecase

import (
	"belajar-api/app/book/repository"
	// borrowRepository "belajar-api/app/borrowhistories/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
)

type IBookUsecase interface {
	GetAllBooks() ([]domain.Books, any)
	GetBookById(id int) (domain.Books, any)
	GetAvailableBook() ([]domain.Books, any)
	CreateBook(bookRequest domain.BookRequest) (domain.Books, any)
	Update(id int, bookRequest domain.BookRequest) (domain.Books, any)
	Delete(id int) (domain.Books, any)
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
			Message: "error occured when get all book",
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
			Message: "error occured when get available book",
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
