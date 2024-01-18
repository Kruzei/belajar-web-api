package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IBookRepository interface {
	FindAll() ([]domain.Books, error)
	FindById(id int) (domain.Books, error)
	CreateBook(book domain.Books) (domain.Books, error)
	Update(book domain.Books) (domain.Books, error)
	Delete(book domain.Books) (domain.Books, error)
}

type BookRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) FindAll() ([]domain.Books, error) {
	var books []domain.Books
	err := r.db.Find(&books).Error
	return books, err
}

func (r *BookRepository) FindById(id int) (domain.Books, error) {
	var book domain.Books

	err := r.db.Find(&book, id).Error

	return book, err
}

func (r *BookRepository) CreateBook(book domain.Books) (domain.Books, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *BookRepository) Update(book domain.Books) (domain.Books, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *BookRepository) Delete(book domain.Books) (domain.Books, error) {
	err := r.db.Delete(&book).Error

	return book, err
}
