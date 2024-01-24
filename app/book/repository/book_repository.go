package repository

import (
	"belajar-api/domain"

	"gorm.io/gorm"
)

type IBookRepository interface {
	FindAll(books *[]domain.Books) (error)
	FindById(book *domain.Books, id int) (error)
	CreateBook(book *domain.Books) (error)
	Update(book *domain.Books) (error)
	Delete(book *domain.Books) (error)
}

type BookRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) FindAll(books *[]domain.Books) (error) {
	err := r.db.Find(books).Error
	return err
}

func (r *BookRepository) FindById(book *domain.Books, id int) (error) {
	err := r.db.Where("id = ?", id).First(&book).Error

	return err
}

func (r *BookRepository) CreateBook(book *domain.Books) (error) {
	tx := r.db.Begin()

	err := r.db.Create(book).Error
	if err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *BookRepository) Update(book *domain.Books) (error) {
	tx := r.db.Begin()

	err := r.db.Save(book).Error
	if err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *BookRepository) Delete(book *domain.Books) (error) {
	tx := r.db.Begin()

	err := r.db.Delete(book).Error
	if err != nil{
		tx.Rollback()
		return err
	}

	tx.Commit()
	return err
}
