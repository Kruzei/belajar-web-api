package domain

import (
	"time"
)

type Books struct {
	ID          int    `json:"id" gorm:"primary key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt       time.Time         `json:"-"`
	UpdatedAt       time.Time         `json:"-"`
	BorrowHistories []BorrowHistories `json:"-" gorm:"foreignKey:book_id;references:id"`
}

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}
