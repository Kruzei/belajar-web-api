package usecase

import (
	"belajar-api/app/book/repository"
	"belajar-api/domain"
)

type IBookUsecase interface {
	FindAllBooks() ([]domain.Books, error)
	FindBookById(id int) (domain.Books, error)
	CreateBook(bookRequest domain.BookRequest) (domain.Books, error)
	Update(id int, bookRequest domain.BookRequest) (domain.Books, error)
	Delete(id int) (domain.Books, error)
}

type BookUsecase struct {
	bookRepository repository.IBookRepository
}

func NewBookUsecase(repository repository.IBookRepository) *BookUsecase {
	return &BookUsecase{repository}
}

func (s *BookUsecase) FindAllBooks() ([]domain.Books, error) {
	return s.bookRepository.FindAll()
}

func (s *BookUsecase) FindBookById(id int) (domain.Books, error) {
	books, err := s.bookRepository.FindById(id)
	return books, err
}

func (s *BookUsecase) CreateBook(bookRequest domain.BookRequest) (domain.Books, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := domain.Books{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.bookRepository.CreateBook(book)

	return newBook, err
}

func (s *BookUsecase) Update(id int, bookRequest domain.BookRequest) (domain.Books, error) {
	book, _ := s.bookRepository.FindById(id)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Rating = int(rating)
	book.Description = bookRequest.Description

	updatedBook, err := s.bookRepository.Update(book)
	return updatedBook, err
}

func (s *BookUsecase) Delete(id int) (domain.Books, error) {
	book, _ := s.bookRepository.FindById(id)
	deletedBook, err := s.bookRepository.Delete(book)

	return deletedBook, err
}
