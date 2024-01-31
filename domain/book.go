package domain

import (
	"time"
)

//	type Books struct {
//		ID          int    `json:"id"`
//		Title       string `json:"title"`
//		Description string `json:"description"`
//		Price       int    `json:"price"`
//		Rating      int    `json:"rating"`
//		//status enum (avaible/notavaible)
//		//user_id (tempat foreign key)
//		CreatedAt time.Time `json:"-"`
//		UpdatedAt time.Time `json:"-"`
//	}
type Books struct {
	ID              int               `json:"id" gorm:"primary key"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	Status          string            `json:"status"`
	Users           []Users           `json:"-" gorm:"many2many:borrowhistories;foreignKey:id;joinForeignKey:book_id;references:id;joinReferences:user_id"`
	CreatedAt       time.Time         `json:"-"`
	UpdatedAt       time.Time         `json:"-"`
	// User      []BorrowHistories `json:"-" gorm:"foreignKey:book_id;references:id"`
}

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}
