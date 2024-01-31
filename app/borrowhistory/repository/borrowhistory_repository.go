package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IBorrowHistory interface {
	GetBorrowedBook(books *[]domain.Books, users *[]domain.Users) error
	GetBorrowHistory(borrowHistories *[]domain.BorrowHistories) error
	GetAllBookById( books *[]domain.Books) error
}

type BorrowHistoryRepository struct {
	db *gorm.DB
}

func NewBorrowHistoryRepository(db *gorm.DB) *BorrowHistoryRepository {
	return &BorrowHistoryRepository{db}
}

func (r *BorrowHistoryRepository) GetBorrowedBook(books *[]domain.Books, users *[]domain.Users) error {
	err := r.db.Model(books).Association("Users").Find(&users)
	return err
}

func (r *BorrowHistoryRepository) GetBorrowHistory(borrowHistories *[]domain.BorrowHistories) error{
	err := r.db.Find(&borrowHistories).Error
	return err
}

func (r *BorrowHistoryRepository) GetAllBookById(books *[]domain.Books) error{
	err := r.db.Joins("join borrowhistories on borrowhistories.book_id = books.id").Find(&books).Error
	return err
}
