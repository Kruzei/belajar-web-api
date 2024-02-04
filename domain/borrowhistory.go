package domain

import "time"

type BorrowHistories struct {
	Id         int       `json:"id" gorm:"primary key"`
	UserId     int       `json:"-" gorm:"user_id"`
	BookId     int       `json:"-" gorm:"book_id"`
	User       Users     `json:"-"`
	Book       Books     `json:"-"`
	BorrowTime time.Time `json:"borrow_time"`
	ReturnTime time.Time `json:"return_time"`
}

func (b *BorrowHistories) TableName() string {
	return "borrowhistories"
}

type BorrowedBookResponse struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BorrowedBookByUserResponse struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BorrowHistoryResponse struct {
	Title      string    `json:"title"`
	BorrowTime time.Time `json:"borrow_time"`
	ReturnTime time.Time `json:"return_time"`
}
