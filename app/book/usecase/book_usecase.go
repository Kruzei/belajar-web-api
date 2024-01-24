package usecase

import (
	"belajar-api/app/book/repository"
	"belajar-api/domain"
	help "belajar-api/helper"
	"errors"
	"net/http"
)

type IBookUsecase interface {
	FindAllBooks() ([]domain.Books, any)
	FindBookById(id int) (domain.Books, any)
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

func (s *BookUsecase) FindAllBooks() ([]domain.Books, any) {
	var books []domain.Books
	err := s.bookRepository.FindAll(&books)
	if err != nil{
		return books, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "Error occured when find all book",
			Err: err,
		}
	}
	
	return books, err
}

func (s *BookUsecase) FindBookById(id int) (domain.Books, any) {
	var book domain.Books 
	err := s.bookRepository.FindById(&book, id)
	if err != nil{
		return domain.Books{}, help.ErrorObject{
			Code: http.StatusNotFound,
			Message: "ID is not exist",
			Err: errors.New("ID NOT FOUND"),
		}
	}

	return book, err
}

func (s *BookUsecase) CreateBook(bookRequest domain.BookRequest) (domain.Books, any) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := domain.Books{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	err := s.bookRepository.CreateBook(&book)
	if err != nil{
		return domain.Books{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "Failed to make book data",
			Err: err,
		}
	}

	return book, nil
}

func (s *BookUsecase) Update(id int, bookRequest domain.BookRequest) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.FindById(&book, id)
	if err != nil{
		return domain.Books{}, help.ErrorObject{
			Code: http.StatusNotFound,
			Message: "ID is not exist",
			Err: errors.New("ID NOT FOUND"),
		}
	}


	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Rating = int(rating)
	book.Description = bookRequest.Description

	err = s.bookRepository.Update(&book)
	if err != nil{
		return domain.Books{}, help.ErrorObject{
			Code: http.StatusInternalServerError,
			Message: "Failed to update book data",
			Err: err,
		}
	}
	return book, err
}

func (s *BookUsecase) Delete(id int) (domain.Books, any) {
	var book domain.Books
	err := s.bookRepository.FindById(&book, id)
	if err != nil{
		return domain.Books{}, help.ErrorObject{
			Code: http.StatusNotFound,
			Message: "ID is not exist",
			Err: errors.New("ID NOT FOUND"),
		}
	}

	err = s.bookRepository.Delete(&book)

	return book, err
}
