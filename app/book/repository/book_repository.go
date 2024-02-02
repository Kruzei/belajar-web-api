package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IBookRepository interface {
	GetAllBooks(books *[]domain.Books) error
	GetBookById(book *domain.Books, id int) error
	GetAvailableBook(books *[]domain.Books) error
	GetBookByCondition(books *[]domain.Books, condition string, value any) error
	GetBorrowedBook(book *domain.BorrowHistories, bookId int) error
	CreateBook(book *domain.Books) error
	Update(book *domain.Books) error
	Delete(book *domain.Books) error
	Borrow(borrow *domain.BorrowHistories) error
	ReturnBook(borrow *domain.BorrowHistories) error
}

type BookRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) GetAllBooks(books *[]domain.Books) error {
	err := r.db.Find(books).Error
	return err
}

func (r *BookRepository) GetBookById(book *domain.Books, id int) error {
	err := r.db.Where("id = ?", id).First(&book).Error
	return err
}

func (r *BookRepository) GetAvailableBook(books *[]domain.Books) error {
	err := r.db.Where("status = ?", "AVAILABLE").Find(&books).Error
	return err
}

func (r *BookRepository) GetBookByCondition(books *[]domain.Books, condition string, value any) error {
	err := r.db.Where(condition, value).Find(&books).Error
	return err
}

func (r *BookRepository) GetBorrowedBook(book *domain.BorrowHistories, bookId int) error {
	err := r.db.Table("borrowhistories").Where("book_id = ?", bookId).Last(&book).Error

	return err
}

func (r *BookRepository) CreateBook(book *domain.Books) error {
	tx := r.db.Begin()

	err := r.db.Create(book).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *BookRepository) Update(book *domain.Books) error {
	tx := r.db.Begin()

	err := r.db.Save(book).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *BookRepository) Delete(book *domain.Books) error {
	tx := r.db.Begin()

	err := r.db.Delete(book).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}

func (r *BookRepository) Borrow(borrow *domain.BorrowHistories) error {
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

func (r *BookRepository) ReturnBook(borrow *domain.BorrowHistories) error {
	tx := r.db.Begin()

	err := r.db.Table("borrowhistories").Save(borrow).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
