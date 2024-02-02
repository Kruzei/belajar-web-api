package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IBorrowHistory interface {
	GetBorrowedHistories(borrowHistories *[]domain.BorrowHistories) error
	GetBorrowHistory(borrowHistories *[]domain.BorrowHistories) error
	GetBorrowedBooks(borrowHistory *[]domain.BorrowHistories) error
}

type BorrowHistoryRepository struct {
	db *gorm.DB
}

func NewBorrowHistoryRepository(db *gorm.DB) *BorrowHistoryRepository {
	return &BorrowHistoryRepository{db}
}

func (r *BorrowHistoryRepository) GetBorrowedHistories(borrowHistories *[]domain.BorrowHistories) error {
	err := r.db.Joins("join books on books.id = borrowhistories.book_id").Find(&borrowHistories).Error
	return err
}

func (r *BorrowHistoryRepository) GetBorrowedBooks(borrowHistory *[]domain.BorrowHistories) error {
	err := r.db.Model(domain.BorrowHistories{}).Preload("User").Preload("Book", "status = ?", "NOT AVAIBLE").Find(borrowHistory).Error
	return err
}

func (r *BorrowHistoryRepository) GetBorrowHistory(borrowHistories *[]domain.BorrowHistories) error {
	err := r.db.Model(domain.BorrowHistories{}).Preload("Book").Find(&borrowHistories).Error
	return err
}
