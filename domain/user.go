package domain

import (
	"time"
)

type UsersRequests struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users struct {
	ID              int               `json:"id" gorm:"primary key"`
	Name            string            `json:"name"`
	Password        string            `json:"password"`
	Email           string            `json:"email" gorm:"unique"`
	Role            string            `json:"role"`
	BorrowHistories []BorrowHistories `json:"-" gorm:"foreignKey:user_id;references:id"`
	CreatedAt       time.Time         `json:"-"`
	UpdatedAt       time.Time         `json:"-"`
}

type UserParam struct {
	ID    int
	Name  string
	Email string
}

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
