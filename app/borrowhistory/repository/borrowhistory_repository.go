package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IBorrowHistory interface {
	GetBorrowedHistories(borrowHistories *[]domain.BorrowHistories) error
	GetBorrowHistory(borrowHistories *[]domain.BorrowHistories) error
	GetBorrowedBooks(borrowHistory *[]domain.BorrowHistories) error
	GetUserBorrowedBook(borrowedBooks *[]domain.BorrowHistories, userId int) error
	Borrow(borrow *domain.BorrowHistories) error
	ReturnBook(borrow *domain.BorrowHistories) error
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

func (r *BorrowHistoryRepository) GetUserBorrowedBook(borrowedBooks *[]domain.BorrowHistories, userId int) error {
	err := r.db.Debug().Model(domain.BorrowHistories{}).Preload("Book").Where("return_time is NULL").Find(&borrowedBooks, "user_id = ?", userId).Error
	return err
}

func (r *BorrowHistoryRepository) Borrow(borrow *domain.BorrowHistories) error {
	tx := r.db.Begin()

	err := r.db.Table("borrowhistories").Create(map[string]interface{}{
		"user_id":     borrow.UserId,
		"book_id":     borrow.BookId,
		"borrow_time": borrow.BorrowTime,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (r *BorrowHistoryRepository) ReturnBook(borrow *domain.BorrowHistories) error {
	tx := r.db.Begin()

	err := r.db.Table("borrowhistories").Save(borrow).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
