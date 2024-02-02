package domain

import (
	"time"
)

// type Users struct {
// 	ID       int    `json:"id" gorm:"primary key"`
// 	Email    string `json:"email" gorm:"unique"`
// 	Password string `json:"password"`
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Gender   string `json:"gender"`
// 	//status enum admin/client
// 	CreatedAt time.Time `json:"-"`
// 	UpdatedAt time.Time `json:"-"`
// }

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
	// BorrowHistories []BorrowHistories `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
